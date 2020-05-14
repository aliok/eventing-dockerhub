package adapter

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"
	"time"
)

const (
	DHHeaderEvent    = "DockerHub-Event"
	DHHeaderDelivery = "DockerHub-Delivery"
)

type envConfig struct {
	// Include the standard adapter.EnvConfig used by all adapters.
	adapter.EnvConfig
	// Port to listen incoming connections
	Port string `envconfig:"PORT" default:"8080"`

}

func NewEnv() adapter.EnvConfigAccessor { return &envConfig{} }

// Adapter converts incoming GitHub webhook events to CloudEvents
type Adapter struct {
	client cloudevents.Client
	source string
	logger   *zap.SugaredLogger
	port string
}

// NewAdapter creates an adapter to convert incoming DockerHub webhook events to CloudEvents and
// then sends them to the specified Sink
func NewAdapter(ctx context.Context, aEnv adapter.EnvConfigAccessor, ceClient cloudevents.Client) adapter.Adapter {
	env := aEnv.(*envConfig) // Will always be our own envConfig type
	logger := logging.FromContext(ctx)
	return &Adapter{
		client:   ceClient,
		logger:   logger,
		port: env.Port,
	}
}

// Start runs the adapter.
// Returns if stopCh is closed or Send() returns an error.
func (a *Adapter) Start(stopCh <-chan struct{}) error {
	done := make(chan bool, 1)

	server := &http.Server{
		Addr:    ":" + a.port,
		Handler: a.newRouter(),
	}

	go gracefulShutdown(server, a.logger, stopCh, done)

	a.logger.Infof("Server is ready to handle requests at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen on %s: %v", server.Addr, err)
	}

	<-done
	a.logger.Infof("Server stopped")
	return nil
}

func gracefulShutdown(server *http.Server, logger *zap.SugaredLogger, stopCh <-chan struct{}, done chan<- bool) {
	<-stopCh
	logger.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v", err)
	}
	close(done)
}

// HandleEvent is invoked whenever an event comes in from GitHub
func (a *Adapter) HandleEvent(payload interface{}, header http.Header) {
	hdr := http.Header(header)
	err := a.handleEvent(payload, hdr)
	if err != nil {
		//log.Printf("unexpected error handling GitHub event: %s", err)
	}
}

func (a *Adapter) newRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO Add more
		//err = ra.handleEvent(payload, r.Header)
		//if err != nil {
		//	ra.logger.Errorf("event handler error: %v", err)
		//	w.WriteHeader(400)
		//	w.Write([]byte(err.Error()))
		//	return
		//}
		//ra.logger.Infof("event processed")
		//w.WriteHeader(202)
		//w.Write([]byte("accepted"))
	})
	return router
}

func (a *Adapter) handleEvent(payload interface{}, hdr http.Header) error {
	dockerHubEventType := hdr.Get("X-" + DHHeaderEvent)
	eventID := hdr.Get("X-" + DHHeaderDelivery)

	event := cloudevents.NewEvent(cloudevents.VersionV1)
	event.SetID(eventID)
	event.SetType(dockerHubEventType)
	event.SetSource(a.source)
	err := event.SetData(cloudevents.ApplicationJSON, payload)
	if err != nil {
		return err
	}

	result := a.client.Send(context.Background(), event)
	if !cloudevents.IsACK(result) {
		return result
	}
	return nil
}
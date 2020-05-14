package main

import (
	"knative.dev/eventing/pkg/adapter/v2"
	myadapter "github.com/tom24d/eventing-dockerhub/pkg/adapter"
)

func main() {
	// TODO impl to read env.
	adapter.Main("dockerhub-source", myadapter.NewEnv, myadapter.NewAdapter)
}

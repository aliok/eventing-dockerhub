package resources

// import (
// 	"fmt"

// 	v1 "k8s.io/api/apps/v1"
// 	corev1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"knative.dev/pkg/kmeta"

// 	"knative.dev/sample-source/pkg/apis/samples/v1alpha1"
// )

// // ReceiveAdapterArgs are the arguments needed to create a Sample Source Receive Adapter.
// // Every field is required.
// type ReceiveAdapterArgs struct {
// 	Image       string
// 	Labels      map[string]string
// 	Source      *v1alpha1.SampleSource
// 	EventSource string
// }

// // MakeReceiveAdapter generates (but does not insert into K8s) the Receive Adapter Deployment for
// // Sample sources.
// func MakeReceiveAdapter(args *ReceiveAdapterArgs) *v1.Deployment {
// 	replicas := int32(1)
// 	return &v1.Deployment{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Namespace: args.Source.Namespace,
// 			Name:      kmeta.ChildName(fmt.Sprintf("samplesource-%s-", args.Source.Name), string(args.Source.GetUID())),
// 			Labels:    args.Labels,
// 			OwnerReferences: []metav1.OwnerReference{
// 				*kmeta.NewControllerRef(args.Source),
// 			},
// 		},
// 		Spec: v1.DeploymentSpec{
// 			Selector: &metav1.LabelSelector{
// 				MatchLabels: args.Labels,
// 			},
// 			Replicas: &replicas,
// 			Template: corev1.PodTemplateSpec{
// 				ObjectMeta: metav1.ObjectMeta{
// 					Labels: args.Labels,
// 				},
// 				Spec: corev1.PodSpec{
// 					ServiceAccountName: args.Source.Spec.ServiceAccountName,
// 					Containers: []corev1.Container{
// 						{
// 							Name:  "receive-adapter",
// 							Image: args.Image,
// 							Env:   makeEnv(args.EventSource, &args.Source.Spec),
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func makeEnv(eventSource string, spec *v1alpha1.SampleSourceSpec) []corev1.EnvVar {
// 	return []corev1.EnvVar{{
// 		Name:  "EVENT_SOURCE",
// 		Value: eventSource,
// 	}, {
// 		Name:  "INTERVAL",
// 		Value: spec.Interval,
// 	}, {
// 		Name: "NAMESPACE",
// 		ValueFrom: &corev1.EnvVarSource{
// 			FieldRef: &corev1.ObjectFieldSelector{
// 				FieldPath: "metadata.namespace",
// 			},
// 		},
// 	}, {
// 		Name: "NAME",
// 		ValueFrom: &corev1.EnvVarSource{
// 			FieldRef: &corev1.ObjectFieldSelector{
// 				FieldPath: "metadata.name",
// 			},
// 		},
// 	}, {
// 		Name:  "METRICS_DOMAIN",
// 		Value: "knative.dev/eventing",
// 	}, {
// 		Name:  "K_METRICS_CONFIG",
// 		Value: "",
// 	}, {
// 		Name:  "K_LOGGING_CONFIG",
// 		Value: "",
// 	}}
// }

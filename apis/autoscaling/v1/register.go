package v1

import "github.com/sourcegraph/k8s"

func init() {
	k8s.Register("autoscaling", "v1", "horizontalpodautoscalers", true, &HorizontalPodAutoscaler{})

	k8s.RegisterList("autoscaling", "v1", "horizontalpodautoscalers", true, &HorizontalPodAutoscalerList{})
}

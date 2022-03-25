package v1

import "github.com/sourcegraph/k8s"

func init() {
	k8s.Register("batch", "v1", "jobs", true, &Job{})

	k8s.RegisterList("batch", "v1", "jobs", true, &JobList{})
}

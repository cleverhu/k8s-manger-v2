package configs

import (
	"k8s-manger-v2/services"
)

type K8sMaps struct {
}

func NewK8sMaps() *K8sMaps {
	return &K8sMaps{}
}

func (*K8sMaps) DeploymentMap() *services.DeploymentMap {
	return &services.DeploymentMap{}
}

func (*K8sMaps) PodMap() *services.PodMap {
	return &services.PodMap{}
}

func (*K8sMaps) RSMap() *services.RSMap {
	return &services.RSMap{}
}

func (*K8sMaps) NSMap() *services.NSMap {
	return &services.NSMap{}
}

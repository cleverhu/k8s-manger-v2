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

func (*K8sMaps) EventMap() *services.EventMap {
	return &services.EventMap{}
}

func (*K8sMaps) IngressMap() *services.IngressMap {
	return &services.IngressMap{}
}

func (*K8sMaps) ServiceMap() *services.ServiceMap {
	return &services.ServiceMap{}
}

func (*K8sMaps) SecretMap() *services.SecretMap {
	return &services.SecretMap{}
}

func (*K8sMaps) ConfigMapMap() *services.ConfigMapMap {
	return &services.ConfigMapMap{}
}

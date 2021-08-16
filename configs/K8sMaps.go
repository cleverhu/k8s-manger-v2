package configs

import "k8s-manger-v2/core"

type K8sMaps struct {
}

func NewK8sMaps() *K8sMaps {
	return &K8sMaps{}
}

func (*K8sMaps) DeploymentMap() *core.DeploymentMap {
	return &core.DeploymentMap{}
}

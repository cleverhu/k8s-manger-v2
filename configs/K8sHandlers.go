package configs

import "k8s-manger-v2/core"

type K8sHandlers struct {
}

func NewK8sHandlers() *K8sHandlers {
	return &K8sHandlers{}
}

func (*K8sHandlers) DepHandler() *core.DepHandler {
	return &core.DepHandler{}
}

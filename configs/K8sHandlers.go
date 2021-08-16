package configs

import (
	"k8s-manger-v2/services"
)

type K8sHandlers struct {
}

func NewK8sHandlers() *K8sHandlers {
	return &K8sHandlers{}
}

func (*K8sHandlers) DepHandler() *services.DepHandler {
	return &services.DepHandler{}
}

func (*K8sHandlers) PodHandler() *services.PodHandler {
	return &services.PodHandler{}
}

func (*K8sHandlers) RSHandler() *services.RSHandler {
	return &services.RSHandler{}
}

func (*K8sHandlers) NSHandler() *services.NSHandler {
	return &services.NSHandler{}
}

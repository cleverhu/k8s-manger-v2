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

func (*K8sHandlers) EventHandler() *services.EventHandler {
	return &services.EventHandler{}
}

func (*K8sHandlers) IngressHandler() *services.IngressHandler {
	return &services.IngressHandler{}
}

func (*K8sHandlers) ServiceHandler() *services.ServiceHandler {
	return &services.ServiceHandler{}
}

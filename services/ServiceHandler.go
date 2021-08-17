package services

import (
	"k8s-manger-v2/wscore"
	corev1 "k8s.io/api/core/v1"
	"log"
)

type ServiceHandler struct {
	ServiceMap     *ServiceMap     `inject:"-"`
	ServiceService *ServiceService `inject:"-"`
}

func (this *ServiceHandler) OnAdd(obj interface{}) {
	service := obj.(*corev1.Service)
	this.ServiceMap.Add(service)
	wscore.ClientMap.SendAll("Services", service.Namespace, this.ServiceService.GetServiceListByNS(service.Namespace))
}
func (this *ServiceHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	service := newObj.(*corev1.Service)
	err := this.ServiceMap.Update(service)
	if err != nil {
		log.Println(err)
	} else {
		wscore.ClientMap.SendAll("Services", service.Namespace, this.ServiceService.GetServiceListByNS(service.Namespace))
	}
}
func (this *ServiceHandler) OnDelete(obj interface{}) {
	service := obj.(*corev1.Service)
	this.ServiceMap.Delete(service)
	wscore.ClientMap.SendAll("Services", service.Namespace, this.ServiceService.GetServiceListByNS(service.Namespace))
}

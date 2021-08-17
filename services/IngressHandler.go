package services

import (
	"k8s-manger-v2/wscore"
	"k8s.io/api/networking/v1beta1"
	"log"
)

type IngressHandler struct {
	IngressMap     *IngressMap     `inject:"-"`
	IngressService *IngressService `inject:"-"`
}

func (this *IngressHandler) OnAdd(obj interface{}) {
	ingress := obj.(*v1beta1.Ingress)
	this.IngressMap.Add(ingress)
	wscore.ClientMap.SendAll("Ingress", ingress.Namespace, this.IngressService.GetIngressListByNS(ingress.Namespace))

}
func (this *IngressHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	ingress := newObj.(*v1beta1.Ingress)
	err := this.IngressMap.Update(ingress)
	if err != nil {
		log.Println(err)
	} else {
		wscore.ClientMap.SendAll("Ingress", ingress.Namespace, this.IngressService.GetIngressListByNS(ingress.Namespace))
	}
}
func (this *IngressHandler) OnDelete(obj interface{}) {
	ingress := obj.(*v1beta1.Ingress)
	this.IngressMap.Delete(ingress)
	wscore.ClientMap.SendAll("Ingress", ingress.Namespace, this.IngressService.GetIngressListByNS(ingress.Namespace))

}

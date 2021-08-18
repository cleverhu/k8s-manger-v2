package services

import (
	corev1 "k8s.io/api/core/v1"
	"log"
)

type NSHandler struct {
	NSMap     *NSMap     `inject:"-"`
	NSService *NSService `inject:"-"`
}

func (this *NSHandler) OnAdd(obj interface{}) {
	this.NSMap.Add(obj.(*corev1.Namespace))
	//fmt.Println(obj.(*corev1.Namespace))

}
func (this *NSHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	//fmt.Println(newObj.(*corev1.Namespace))
	err := this.NSMap.Update(newObj.(*corev1.Namespace))
	if err != nil {
		log.Println(err)
	}
}
func (this *NSHandler) OnDelete(obj interface{}) {
	this.NSMap.Delete(obj.(*corev1.Namespace))
}

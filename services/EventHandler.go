package services

import (
	corev1 "k8s.io/api/core/v1"
	"log"
)

type EventHandler struct {
	EventMap *EventMap `inject:"-"`
}

func (this *EventHandler) OnAdd(obj interface{}) {
	this.EventMap.Add(obj.(*corev1.Event))
}
func (this *EventHandler) OnUpdate(oldObj interface{}, newObj interface{}) {

	err := this.EventMap.Update(newObj.(*corev1.Event))
	if err != nil {
		log.Println(err)
	}
}
func (this *EventHandler) OnDelete(obj interface{}) {
	this.EventMap.Delete(obj.(*corev1.Event))
}

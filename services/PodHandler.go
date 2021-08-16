package services

import (
	corev1 "k8s.io/api/core/v1"
	"log"
)

type PodHandler struct {
	PodMap *PodMap `inject:"-"`
}

func (this *PodHandler) OnAdd(obj interface{}) {
	this.PodMap.Add(obj.(*corev1.Pod))
}
func (this *PodHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	err := this.PodMap.Update(newObj.(*corev1.Pod))
	if err != nil {
		log.Println(err)
	}
}
func (this *PodHandler) OnDelete(obj interface{}) {
	this.PodMap.Delete(obj.(*corev1.Pod))
}

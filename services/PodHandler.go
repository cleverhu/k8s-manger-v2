package services

import (
	"k8s-manger-v2/wscore"
	corev1 "k8s.io/api/core/v1"
	"log"
)

type PodHandler struct {
	PodMap     *PodMap     `inject:"-"`
	PodService *PodService `inject:"-"`
}

func (this *PodHandler) OnAdd(obj interface{}) {
	pod := obj.(*corev1.Pod)
	this.PodMap.Add(pod)

	wscore.ClientMap.SendAll("Pods", pod.Namespace, this.PodService.GetPodsListByNS(pod.Namespace))
}
func (this *PodHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	pod := newObj.(*corev1.Pod)
	err := this.PodMap.Update(pod)
	if err != nil {
		log.Println(err)
	} else {
		wscore.ClientMap.SendAll("Pods", pod.Namespace, this.PodService.GetPodsListByNS(pod.Namespace))

	}
}
func (this *PodHandler) OnDelete(obj interface{}) {
	pod := obj.(*corev1.Pod)

	this.PodMap.Delete(pod)
	wscore.ClientMap.SendAll("Pods", pod.Namespace, this.PodService.GetPodsListByNS(pod.Namespace))
}

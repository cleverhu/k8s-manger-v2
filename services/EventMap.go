package services

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sync"
)

type EventMap struct {
	CommonService *CommonService `inject:"-"`
	data          sync.Map       //key ns value []*corev1.Pod
}

func (this *EventMap) Add(event *corev1.Event) {
	key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)

	this.data.Store(key, event.Message)
}

func (this *EventMap) Delete(event *corev1.Event) {
	key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
	this.data.Delete(key)
}

func (this *EventMap) Update(event *corev1.Event) error {
	key := fmt.Sprintf("%s_%s_%s", event.Namespace, event.InvolvedObject.Kind, event.InvolvedObject.Name)
	this.data.Store(key, event.Message)
	return nil
}

func (this *EventMap) GetMessage(ns, kind, name string) string {
	key := fmt.Sprintf("%s_%s_%s", ns, kind, name)
	value, ok := this.data.Load(key)
	if ok {
		return value.(string)
	}
	return ""
}

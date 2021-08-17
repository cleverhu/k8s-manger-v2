package services

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type ServiceMap struct {
	data sync.Map //key ns value []*corev1.Pod
}

func (this *ServiceMap) Add(pod *corev1.Service) {
	key := pod.Namespace
	if value, ok := this.data.Load(key); ok {
		value = append(value.([]*corev1.Service), pod)
		this.data.Store(key, value)
	} else {
		this.data.Store(key, []*corev1.Service{pod})
	}
}

func (this *ServiceMap) Delete(pod *corev1.Service) {
	key := pod.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, p := range value.([]*corev1.Service) {
			if p.Name == pod.Name {
				value = append(value.([]*corev1.Service)[0:index], value.([]*corev1.Service)[index+1:]...)
				this.data.Store(key, value)
				return
			}
		}
	}
}

func (this *ServiceMap) Update(pod *corev1.Service) error {
	key := pod.Namespace
	if value, ok := this.data.Load(key); ok {

		for index, p := range value.([]*corev1.Service) {
			if p.Name == pod.Name {
				value.([]*corev1.Service)[index] = pod
				this.data.Store(key, value)
				return nil
			}
		}
	}
	return fmt.Errorf("pod-%s not found", pod.Name)
}

func (this *ServiceMap) ListByNS(ns string) ([]*corev1.Service, error) {
	if ns != "all-namespaces" {
		if list, ok := this.data.Load(ns); ok {
			return list.([]*corev1.Service), nil
		}
	} else {
		ret := make([]*corev1.Service, 0)
		this.data.Range(func(key, value interface{}) bool {
			for _, dep := range value.([]*corev1.Service) {
				ret = append(ret, dep)
			}
			sort.Slice(ret, func(i, j int) bool {
				if ret[i].Namespace == ret[j].Namespace {
					return ret[i].Name < ret[j].Name
				} else {
					return ret[i].Namespace < ret[j].Namespace
				}
			})
			return true
		})
		return ret, nil
	}

	return nil, nil
	//return nil, fmt.Errorf("pods not found")
}

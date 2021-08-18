package services

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type ConfigMapMap struct {
	data sync.Map //key ns value []*corev1.Pod
}

func (this *ConfigMapMap) Add(cm *corev1.ConfigMap) {
	key := cm.Namespace
	if value, ok := this.data.Load(key); ok {
		value = append(value.([]*corev1.ConfigMap), cm)
		this.data.Store(key, value)
	} else {
		this.data.Store(key, []*corev1.ConfigMap{cm})
	}
}

func (this *ConfigMapMap) Delete(cm *corev1.ConfigMap) {
	key := cm.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, p := range value.([]*corev1.ConfigMap) {
			if p.Name == cm.Name {
				value = append(value.([]*corev1.ConfigMap)[0:index], value.([]*corev1.ConfigMap)[index+1:]...)
				this.data.Store(key, value)
				return
			}
		}
	}
}

func (this *ConfigMapMap) Update(cm *corev1.ConfigMap) error {
	key := cm.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, p := range value.([]*corev1.ConfigMap) {
			if p.Name == cm.Name {
				value.([]*corev1.ConfigMap)[index] = cm
				this.data.Store(key, value)
				return nil
			}
		}
	}
	return fmt.Errorf("configmap-%s not found", cm.Name)
}

func (this *ConfigMapMap) ListByNS(ns string) []*corev1.ConfigMap {

	if ns != "all-namespaces" {
		if list, ok := this.data.Load(ns); ok {
			return list.([]*corev1.ConfigMap)
		}
	} else {
		ret := make([]*corev1.ConfigMap, 0)
		this.data.Range(func(key, value interface{}) bool {
			for _, secret := range value.([]*corev1.ConfigMap) {
				ret = append(ret, secret)
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
		return ret
	}

	return nil
}

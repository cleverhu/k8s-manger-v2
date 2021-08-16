package services

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sync"
)

type NSMap struct {
	CommonService *CommonService `inject:"-"`
	data          sync.Map       //key ns value []*corev1.NameSpace
}

func (this *NSMap) Add(ns *corev1.Namespace) {
	key := ns.Name
	//fmt.Println(key)
	if value, ok := this.data.Load(key); ok {
		value = append(value.([]*corev1.Namespace), ns)
		this.data.Store(key, value)
	} else {
		this.data.Store(key, []*corev1.Namespace{ns})
	}
}

func (this *NSMap) Delete(ns *corev1.Namespace) {
	key := ns.Name
	if value, ok := this.data.Load(key); ok {
		for index, p := range value.([]*corev1.Namespace) {
			if p.Name == ns.Name {
				value = append(value.([]*corev1.Namespace)[0:index], value.([]*corev1.Namespace)[index+1:]...)
				this.data.Store(key, value)
				return
			}
		}
	}
}

func (this *NSMap) Update(ns *corev1.Namespace) error {
	key := ns.Name
	if value, ok := this.data.Load(key); ok {

		for index, p := range value.([]*corev1.Namespace) {
			if p.Name == ns.Name {
				value.([]*corev1.Namespace)[index] = ns
				this.data.Store(key, value)
				return nil
			}
		}
	}

	return fmt.Errorf("ns:%s not found", ns.Name)
}

func (this *PodMap) ListAll() ([]*corev1.Pod, error) {
	return nil, fmt.Errorf("ns not found")
}

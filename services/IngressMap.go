package services

import (
	"k8s.io/api/networking/v1beta1"
	"sort"

	"sync"
)

type IngressMap struct {
	CommonService *CommonService `inject:"-"`
	data          sync.Map       //key ns value []*corev1.Pod
}

func (this *IngressMap) Add(ingress *v1beta1.Ingress) {
	key := ingress.Namespace
	if value, ok := this.data.Load(key); ok {
		value = append(value.([]*v1beta1.Ingress), ingress)
		this.data.Store(key, value)
	} else {
		this.data.Store(key, []*v1beta1.Ingress{ingress})
	}
}

func (this *IngressMap) Delete(ingress *v1beta1.Ingress) {
	key := ingress.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, p := range value.([]*v1beta1.Ingress) {
			if p.Name == ingress.Name {
				value = append(value.([]*v1beta1.Ingress)[0:index], value.([]*v1beta1.Ingress)[index+1:]...)
				this.data.Store(key, value)
				return
			}
		}
	}
}

func (this *IngressMap) Update(ingress *v1beta1.Ingress) error {
	key := ingress.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, p := range value.([]*v1beta1.Ingress) {
			if p.Name == ingress.Name {
				value.([]*v1beta1.Ingress)[index] = ingress
				this.data.Store(key, value)
				return nil
			}
		}
	}
	return nil
}

func (this *IngressMap) ListByNS(ns string) ([]*v1beta1.Ingress, error) {
	if ns != "all-namespaces" {
		if list, ok := this.data.Load(ns); ok {
			return list.([]*v1beta1.Ingress), nil
		}
	} else {
		ret := make([]*v1beta1.Ingress, 0)
		this.data.Range(func(key, value interface{}) bool {
			for _, ingress := range value.([]*v1beta1.Ingress) {
				ret = append(ret, ingress)
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

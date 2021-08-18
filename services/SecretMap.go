package services

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type SecretMap struct {
	data sync.Map //key ns value []*corev1.Pod
}

func (this *SecretMap) Add(secret *corev1.Secret) {
	key := secret.Namespace
	if value, ok := this.data.Load(key); ok {
		value = append(value.([]*corev1.Secret), secret)
		this.data.Store(key, value)
	} else {
		this.data.Store(key, []*corev1.Secret{secret})
	}
}

func (this *SecretMap) Delete(secret *corev1.Secret) {
	key := secret.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, p := range value.([]*corev1.Secret) {
			if p.Name == secret.Name {
				value = append(value.([]*corev1.Secret)[0:index], value.([]*corev1.Secret)[index+1:]...)
				this.data.Store(key, value)
				return
			}
		}
	}
}

func (this *SecretMap) Update(secret *corev1.Secret) error {
	key := secret.Namespace
	if value, ok := this.data.Load(key); ok {

		for index, p := range value.([]*corev1.Secret) {
			if p.Name == secret.Name {
				value.([]*corev1.Secret)[index] = secret
				this.data.Store(key, value)
				return nil
			}
		}
	}
	return fmt.Errorf("secret-%s not found", secret.Name)
}

func (this *SecretMap) ListByNS(ns string) ([]*corev1.Secret, error) {

	if ns != "all-namespaces" {
		if list, ok := this.data.Load(ns); ok {
			return list.([]*corev1.Secret), nil
		}
	} else {
		ret := make([]*corev1.Secret, 0)
		this.data.Range(func(key, value interface{}) bool {
			for _, secret := range value.([]*corev1.Secret) {
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
		return ret, nil
	}

	return nil, nil
	//return nil, fmt.Errorf("pods not found")
}

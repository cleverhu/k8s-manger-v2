package core

import (
	"errors"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"sync"
)

type RSMap struct {
	data sync.Map //key ns value []*corev1.Pod
}

func (this *RSMap) Add(rs *v1.ReplicaSet) {
	key := rs.Namespace
	if value, ok := this.data.Load(key); ok {
		value = append(value.([]*v1.ReplicaSet), rs)
		this.data.Store(key, value)
	} else {
		this.data.Store(key, []*v1.ReplicaSet{rs})
	}
}

func (this *RSMap) Delete(rs *v1.ReplicaSet) {
	key := rs.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, r := range value.([]*v1.ReplicaSet) {
			if r.Name == rs.Name {
				value = append(value.([]*v1.ReplicaSet)[0:index], value.([]*v1.ReplicaSet)[index+1:]...)
				this.data.Store(key, value)
				return
			}
		}
	}
}

func (this *RSMap) Update(rs *v1.ReplicaSet) error {
	key := rs.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, r := range value.([]*v1.ReplicaSet) {
			if r.Name == rs.Name {
				value.([]*v1.ReplicaSet)[index] = rs
				this.data.Store(key, value)
				return nil
			}
		}
	}

	return fmt.Errorf("rs-%s not found", rs.Name)
}

func (this *RSMap) ListByNS(ns string) ([]*v1.ReplicaSet, error) {
	if list, ok := this.data.Load(ns); ok {
		return list.([]*v1.ReplicaSet), nil
	}
	return nil, errors.New("rs record not found")
}

func (this *RSMap) GetRsLabelsByDeployment(deploy *v1.Deployment) ([]map[string]string, error) {
	rs, err := this.ListByNS(deploy.Namespace)
	if err != nil {
		return nil, err
	}
	ret := make([]map[string]string, 0)
	for _, item := range rs {
		if item.Annotations["deployment.kubernetes.io/revision"] != deploy.Annotations["deployment.kubernetes.io/revision"] {
			continue
		}
		for _, v := range item.OwnerReferences {
			if v.Name == deploy.Name {
				ret = append(ret, item.Labels)
				break
			}
		}
	}
	return ret, nil
}

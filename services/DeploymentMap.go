package services

import (
	"errors"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"sort"
	"sync"
)

type DeploymentMap struct {
	data sync.Map //key ns value []*v1.Deployment
}

func (this *DeploymentMap) Add(deploy *v1.Deployment) {
	key := deploy.Namespace
	if value, ok := this.data.Load(key); ok {
		value = append(value.([]*v1.Deployment), deploy)
		this.data.Store(key, value)
	} else {
		this.data.Store(key, []*v1.Deployment{deploy})
	}
}

func (this *DeploymentMap) Delete(deploy *v1.Deployment) {
	key := deploy.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, dep := range value.([]*v1.Deployment) {
			if dep.Name == deploy.Name {
				value = append(value.([]*v1.Deployment)[0:index], value.([]*v1.Deployment)[index+1:]...)
				this.data.Store(key, value)
				return
			}
		}
	}
}

func (this *DeploymentMap) Update(deploy *v1.Deployment) error {
	key := deploy.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, dep := range value.([]*v1.Deployment) {
			if dep.Name == deploy.Name {
				value.([]*v1.Deployment)[index] = deploy
				this.data.Store(key, value)
				return nil
			}
		}
	}

	return fmt.Errorf("deployment-%s not found", deploy.Name)
}

func (this *DeploymentMap) ListByNS(ns string) ([]*v1.Deployment, error) {
	if ns != "all-namespaces" {
		if list, ok := this.data.Load(ns); ok {
			return list.([]*v1.Deployment), nil
		}
	} else {
		ret := make([]*v1.Deployment, 0)
		this.data.Range(func(key, value interface{}) bool {
			for _, dep := range value.([]*v1.Deployment) {
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

	return nil, errors.New("deployments record not found")
}

func (this *DeploymentMap) Get(ns, name string) (*v1.Deployment, error) {
	deps, err := this.ListByNS(ns)
	if err != nil {
		return nil, err
	}
	for _, dep := range deps {
		if dep.Name == name {
			return dep, nil
		}
	}
	return nil, errors.New("deployment record not found")
}

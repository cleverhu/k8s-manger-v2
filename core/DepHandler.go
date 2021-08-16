package core

import (
	v1 "k8s.io/api/apps/v1"
	"log"
)

type DepHandler struct {
	DepMap *DeploymentMap `inject:"-"`
}

func (this *DepHandler) OnAdd(obj interface{}) {
	dep := obj.(*v1.Deployment)
	this.DepMap.Add(dep)
}

func (this *DepHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	err := this.DepMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	}
}

func (this *DepHandler) OnDelete(obj interface{}) {
	this.DepMap.Delete(obj.(*v1.Deployment))
}

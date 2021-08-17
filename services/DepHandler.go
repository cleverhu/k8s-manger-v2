package services

import (
	"k8s-manger-v2/wscore"
	v1 "k8s.io/api/apps/v1"
	"log"
)

type DepHandler struct {
	DepMap     *DeploymentMap     `inject:"-"`
	DepService *DeploymentService `inject:"-"`
}

func (this *DepHandler) OnAdd(obj interface{}) {

	dep := obj.(*v1.Deployment)
	//fmt.Println(dep.Namespace)
	this.DepMap.Add(dep)
	wscore.ClientMap.SendAll("Deployments", dep.Namespace, this.DepService.ListAll(dep.Namespace))
}

func (this *DepHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	dep := newObj.(*v1.Deployment)

	err := this.DepMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	} else {
		wscore.ClientMap.SendAll("Deployments", dep.Namespace, this.DepService.ListAll(dep.Namespace))
	}
}

func (this *DepHandler) OnDelete(obj interface{}) {
	dep := obj.(*v1.Deployment)
	this.DepMap.Delete(obj.(*v1.Deployment))
	wscore.ClientMap.SendAll("Deployments", dep.Namespace, this.DepService.ListAll(dep.Namespace))
}

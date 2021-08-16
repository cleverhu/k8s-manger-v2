package core

import (
	v1 "k8s.io/api/apps/v1"
	"log"
)

type RSHandler struct {
	RSMap *RSMap `inject:"-"`
}

func (this *RSHandler) OnAdd(obj interface{}) {
	this.RSMap.Add(obj.(*v1.ReplicaSet))
}
func (this *RSHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	err := this.RSMap.Update(newObj.(*v1.ReplicaSet))
	if err != nil {
		log.Println(err)
	}
}
func (this *RSHandler) OnDelete(obj interface{}) {
	this.RSMap.Delete(obj.(*v1.ReplicaSet))
}

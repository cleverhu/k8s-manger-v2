package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
)

type UserCtl struct {
	K8sCli *kubernetes.Clientset
}

func NewUserCtl() *UserCtl {
	return &UserCtl{}
}

func (*UserCtl) Name() string {
	return "DeploymentCtl"
}

func (this *UserCtl) Index(c *gin.Context) goft.Json {
	return gin.H{"message": "this is index"}
}

func (this *UserCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", this.Index)
}

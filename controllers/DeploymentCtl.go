package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/services"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtl struct {
	K8sCli     *kubernetes.Clientset       `inject:"-"`
	DepService *services.DeploymentService `inject:"-"`
}

func NewUserCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (*DeploymentCtl) Name() string {
	return "DeploymentCtl"
}

func (this *DeploymentCtl) ListAll(c *gin.Context) goft.Json {
	return this.DepService.ListAll("default")
}

func (this *DeploymentCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", this.ListAll)
}

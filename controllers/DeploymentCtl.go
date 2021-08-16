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

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (*DeploymentCtl) Name() string {
	return "DeploymentCtl"
}

func (this *DeploymentCtl) ListAll(c *gin.Context) goft.Json {
	return this.DepService.ListAll("default")
}

func (this *DeploymentCtl) Detail(c *gin.Context) goft.Json {
	return this.DepService.Detail("default", "ngx1")
}

func (this *DeploymentCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", this.ListAll).
		Handle("GET", "/detail", this.Detail)
}

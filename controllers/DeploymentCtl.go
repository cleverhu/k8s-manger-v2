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
	ns := c.Query("namespace")
	if ns == "undefined" || ns == "" {
		ns = "all-namespaces"
	}
	return gin.H{"code": 20000, "data": gin.H{"ns": ns, "data": this.DepService.ListAll(ns)}}

}

func (this *DeploymentCtl) Detail(c *gin.Context) goft.Json {
	return this.DepService.Detail("default", "ngx1")
}

func (this *DeploymentCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/deployments", this.ListAll).
		Handle("GET", "/detail", this.Detail)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/services"
	"k8s.io/client-go/kubernetes"
)

type PodCtl struct {
	K8sCli     *kubernetes.Clientset       `inject:"-"`
	DepService *services.DeploymentService `inject:"-"`
	PodService *services.PodService        `inject:"-"`
}

func NewPodCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (*PodCtl) Name() string {
	return "DeploymentCtl"
}

func (this *PodCtl) ListAll(c *gin.Context) goft.Json {
	return ""
}

func (this *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", this.ListAll)
}

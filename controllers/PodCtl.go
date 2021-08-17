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

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (*PodCtl) Name() string {
	return "DeploymentCtl"
}

func (this *PodCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.Query("namespace")
	if ns == "undefined" || ns == "" {
		ns = "all-namespaces"
	}
	return gin.H{"code": 20000, "data": gin.H{"ns": ns,
		"data": this.PodService.GetPodsListByNS(ns)}}

	//return gin.H{
	//	"code": 20000,
	//	"data": this.PodService.GetPodsListByNS(c.DefaultQuery("namespace", "default")),
	//}
}

func (this *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", this.ListAll)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/services"
	"k8s.io/client-go/kubernetes"
)

type NSCtl struct {
	K8sCli    *kubernetes.Clientset `inject:"-"`
	NSService *services.NSService   `inject:"-"`
}

func NewNSCtl() *NSCtl {
	return &NSCtl{}
}

func (*NSCtl) Name() string {
	return "NSCtl"
}

func (this *NSCtl) ListAll(c *gin.Context) goft.Json {
	return gin.H{"code": 20000, "data": this.NSService.ListAll()}
}

func (this *NSCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ns", this.ListAll)
}

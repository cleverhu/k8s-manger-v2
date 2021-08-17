package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/services"
)

type ServiceCtl struct {
	ServiceService *services.ServiceService `inject:"-"`
}

func NewServiceCtl() *ServiceCtl {
	return &ServiceCtl{}
}

func (*ServiceCtl) Name() string {
	return "SvcCtl"
}

func (this *ServiceCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.Query("namespace")
	if ns == "undefined" || ns == "" {
		ns = "all-namespaces"
	}
	fmt.Println(ns)
	return gin.H{"code": 20000, "data": this.ServiceService.GetServiceListByNS(ns)}
}

func (this *ServiceCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/svc", this.ListAll)
}

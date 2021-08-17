package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/models"
	"k8s-manger-v2/services"
)

type IngressCtl struct {
	IngressService *services.IngressService `inject:"-"`
}

func NewIngressCtl() *IngressCtl {
	return &IngressCtl{}
}

func (*IngressCtl) Name() string {
	return "IngressCtl"
}

func (this *IngressCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.Query("namespace")
	if ns == "undefined" || ns == "" {
		ns = "all-namespaces"
	}
	return gin.H{"code": 20000, "data": gin.H{"ns": ns,
		"data": this.IngressService.GetIngressListByNS(ns)}}
}

func (this *IngressCtl) CreateIngress(c *gin.Context) goft.Json {
	postModel := &models.IngressPost{}
	goft.Error(c.BindJSON(postModel))
	return gin.H{
		"code": 20000,
		"data": postModel,
	}
}

func (this *IngressCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ingress", this.ListAll).
		Handle("POST", "/ingress", this.CreateIngress)
}

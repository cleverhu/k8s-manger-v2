package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/models"
	"k8s-manger-v2/services"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type IngressCtl struct {
	Client         *kubernetes.Clientset    `inject:"-"`
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
	goft.Error(this.IngressService.PostIngress(postModel))
	return gin.H{
		"code": 20000,
		"data": "ok",
	}
}

//DELETE /ingress?ns=xx&name=xx
func (this *IngressCtl) RmIngress(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("namespace", "default")
	name := c.DefaultQuery("name", "")
	goft.Error(this.Client.NetworkingV1beta1().Ingresses(ns).
		Delete(c, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (this *IngressCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ingress", this.ListAll).
		Handle("POST", "/ingress", this.CreateIngress).
		Handle("DELETE", "/ingress", this.RmIngress)
}

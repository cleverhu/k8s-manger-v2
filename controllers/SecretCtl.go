package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/models"
	"k8s-manger-v2/services"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretCtl struct {
	Client        *kubernetes.Clientset   `inject:"-"`
	SecretService *services.SecretService `inject:"-"`
}

func NewSecretCtl() *SecretCtl {
	return &SecretCtl{}
}

func (*SecretCtl) Name() string {
	return "SecretCtl"
}

func (this *SecretCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.Query("namespace")
	if ns == "undefined" || ns == "" {
		ns = "all-namespaces"
	}
	return gin.H{"code": 20000, "data": gin.H{"ns": ns,
		"data": this.SecretService.GetSecretsListByNS(ns)}}
}

func (this *SecretCtl) Detail(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	//goft.Error()
	return gin.H{"code": 20000, "data": this.SecretService.GetSecret(ns, name)}
}

func (this *SecretCtl) RmSecret(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("namespace", "default")
	name := c.DefaultQuery("name", "")
	goft.Error(this.Client.CoreV1().Secrets(ns).Delete(c, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (this *SecretCtl) CreateSecret(c *gin.Context) goft.Json {
	postModel := &models.PostSecretModel{}
	goft.Error(c.BindJSON(postModel))
	goft.Error(this.SecretService.PostSecret(postModel))
	return gin.H{
		"code": 20000,
		"data": "ok",
	}
}

func (this *SecretCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/secrets", this.ListAll).
		Handle("GET", "/secrets/:ns/:name", this.Detail).
		Handle("DELETE", "/secrets", this.RmSecret).
		Handle("POST", "/secrets", this.CreateSecret)
}

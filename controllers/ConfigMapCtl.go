package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/models"
	"k8s-manger-v2/services"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapCtl struct {
	Client           *kubernetes.Clientset      `inject:"-"`
	ConfigMapService *services.ConfigMapService `inject:"-"`
}

func NewConfigMapCtl() *ConfigMapCtl {
	return &ConfigMapCtl{}
}

func (*ConfigMapCtl) Name() string {
	return "SecretCtl"
}

func (this *ConfigMapCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.Query("namespace")
	if ns == "undefined" || ns == "" {
		ns = "all-namespaces"
	}
	return gin.H{"code": 20000, "data": gin.H{"ns": ns,
		"data": this.ConfigMapService.GetConfigMapsListByNS(ns)}}
}

func (this *ConfigMapCtl) Detail(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	//goft.Error()
	return gin.H{"code": 20000, "data": this.ConfigMapService.GetConfigMap(ns, name)}
}

func (this *ConfigMapCtl) RmConfigMap(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("namespace", "default")
	name := c.DefaultQuery("name", "")
	goft.Error(this.Client.CoreV1().ConfigMaps(ns).Delete(c, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "OK",
	}
}

func (this *ConfigMapCtl) CreateConfigMap(c *gin.Context) goft.Json {
	postModel := &models.PostConfigMapModel{}
	goft.Error(c.BindJSON(postModel))
	goft.Error(this.ConfigMapService.PostConfigMap(postModel))
	return gin.H{
		"code": 20000,
		"data": "ok",
	}
}

func (this *ConfigMapCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/configmaps", this.ListAll).
		Handle("GET", "/configmaps/:ns/:name", this.Detail).
		Handle("DELETE", "/configmaps", this.RmConfigMap).
		Handle("POST", "/configmaps", this.CreateConfigMap)
}

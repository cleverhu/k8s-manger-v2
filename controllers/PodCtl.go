package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"io"
	"k8s-manger-v2/services"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net/http"
	"time"
)

type PodCtl struct {
	K8sCli     *kubernetes.Clientset       `inject:"-"`
	DepService *services.DeploymentService `inject:"-"`
	PodService *services.PodService        `inject:"-"`
	Cfg        *rest.Config                `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}

func (*PodCtl) Name() string {
	return "PodCtl"
}

func (this *PodCtl) ListAll(c *gin.Context) goft.Json {
	ns := c.Query("ns")
	if ns == "undefined" || ns == "" {
		ns = "all-namespaces"
	}
	return gin.H{"code": 20000, "data": gin.H{"ns": ns,
		"data": this.PodService.GetPodsListByNS(ns)}}
}

func (this *PodCtl) GetLogs(c *gin.Context) (v goft.Void) {
	ns := c.Query("ns")
	//podName := c.Query("podName")
	//containerName := c.Query("cName")
	podName := c.Query("podName")
	containerName := c.Query("cName")
	var tailLine int64 = 100
	opt := &v1.PodLogOptions{Follow: true, Container: containerName, TailLines: &tailLine}
	fmt.Println(ns, podName, containerName)
	cc, _ := context.WithTimeout(c, time.Minute*30) //设置半小时超时时间。否则会造成内存泄露
	req := this.K8sCli.CoreV1().Pods(ns).GetLogs(podName, opt)
	reader, err := req.Stream(cc)
	goft.Error(err)
	defer reader.Close()
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf) // 如果 当前日志 读完了。 会阻塞

		if err != nil && err != io.EOF { //一旦超时 会进入 这个程序 ,,此时一定要break 掉
			break
		}
		w, err := c.Writer.Write([]byte(string(buf[0:n])))
		if w == 0 || err != nil {
			break
		}
		c.Writer.(http.Flusher).Flush()
	}

	return
}

func (this *PodCtl) GetContainers(c *gin.Context) goft.Json {
	return gin.H{"code": 20000, "data": this.PodService.GetContainers(c.Query("ns"), c.Query("podName"))}
}

func (this *PodCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/pods", this.ListAll).
		Handle("GET", "/pods/logs", this.GetLogs).
		Handle("GET", "/pods/containers", this.GetContainers)
}

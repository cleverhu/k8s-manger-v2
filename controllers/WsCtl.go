package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/helpers"
	"k8s-manger-v2/wscore"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"log"
)

//@Controller
type WsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
	Config *rest.Config          `inject:"-"`
}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func (this *WsCtl) Connect(c *gin.Context) string {
	client, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil) //升级
	if err != nil {
		log.Println(err)
		return err.Error()
	} else {
		wscore.ClientMap.Store(client)
		return "success"
	}

}

func (this *WsCtl) PodConnect(c *gin.Context) (v goft.Void) {
	wsClient, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	shellClient := wscore.NewWsShellClient(wsClient)
	err = helpers.HandleCommand(this.Client, this.Config, []string{"sh"}).
		Stream(remotecommand.StreamOptions{
			Stdin:  shellClient,
			Stdout: shellClient,
			Stderr: shellClient,
			Tty:    true,
		})
	return
}

func (this *WsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ws", this.Connect).
		Handle("GET", "/podws", this.Connect)
}
func (this *WsCtl) Name() string {
	return "WsCtl"
}

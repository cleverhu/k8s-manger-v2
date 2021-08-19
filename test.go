package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s-manger-v2/helpers"
	"k8s-manger-v2/wscore"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
)

func main() {
	cfg, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Insecure = true
	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("hello")
		wsClient, err := wscore.Upgrader.Upgrade(c.Writer, c.Request, nil)
		fmt.Println(err)
		if err != nil {
			return
		}
		shellClient := wscore.NewWsShellClient(wsClient)
		err = helpers.HandleCommand(client, cfg, []string{"sh"}).
			Stream(remotecommand.StreamOptions{
				Stdin:  shellClient,
				Stdout: shellClient,
				Stderr: shellClient,
				Tty:    true,
			})
		if err != nil {
			log.Println(err)
		}

	})
	r.Run(":8080")
}

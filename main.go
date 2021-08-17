package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/configs"
	"k8s-manger-v2/controllers"
	"net/http"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}

func main() {
	server := goft.Ignite(cors()).
		Config(
			configs.NewK8sHandlers(),
			configs.NewK8sConfig(),
			configs.NewK8sMaps(),
			configs.NewServiceConfig(),
		).
		Mount("v1",
			controllers.NewDeploymentCtl(),
			controllers.NewPodCtl(),
			controllers.NewUserCtl(),
			controllers.NewWsCtl(),
			controllers.NewNSCtl(),
		).
		Attach(
		//middlewares.NewCorsMiddleware()
		)
	//server.GET("/dashboard/*filepath", func(c *gin.Context) {
	//	http.FileServer(FS(false)).ServeHTTP(c.Writer, c.Request)
	//})
	server.Launch()
}

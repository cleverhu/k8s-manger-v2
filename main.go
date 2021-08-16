package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/configs"
	"k8s-manger-v2/controllers"
)

func main() {
	goft.Ignite().
		Config(
			configs.NewK8sHandlers(),
			configs.NewK8sConfig(),
			configs.NewK8sMaps(),
			configs.NewServiceConfig(),
		).
		Mount("", controllers.NewDeploymentCtl()).
		Launch()
}

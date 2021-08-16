package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/controllers"
)

func main() {
	goft.Ignite().
		Mount("",controllers.NewUserCtl()).
		Launch()
}

package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
	fmt.Println(client)
}

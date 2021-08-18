package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	pods, _ := client.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})
	fmt.Println(pods.Items[0].Name)
}

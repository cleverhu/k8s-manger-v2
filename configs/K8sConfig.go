package configs

import (
	"k8s-manger-v2/services"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"sync"
)

var K8sClient *kubernetes.Clientset
var k8sClientInitOnce sync.Once

type K8sConfig struct {
	DepHandler *services.DepHandler `inject:"-"`
	PodHandler *services.PodHandler `inject:"-"`
	RSHandler  *services.RSHandler  `inject:"-"`
	NSHandler  *services.NSHandler  `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (this *K8sConfig) K8sClient() *kubernetes.Clientset {
	k8sClientInitOnce.Do(func() {
		config := &rest.Config{Host: "http://47.101.175.193:8009"}
		clientSet, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		K8sClient = clientSet
	})

	return K8sClient
}

func (this *K8sConfig) Informer() informers.SharedInformerFactory {
	factory := informers.NewSharedInformerFactory(this.K8sClient(), 0)
	depInformer := factory.Apps().V1().Deployments().Informer()
	depInformer.AddEventHandler(this.DepHandler)

	podInformer := factory.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(this.PodHandler)

	rsInformer := factory.Apps().V1().ReplicaSets().Informer()
	rsInformer.AddEventHandler(this.RSHandler)

	nsInformer := factory.Core().V1().Namespaces().Informer()
	nsInformer.AddEventHandler(this.NSHandler)

	factory.Start(wait.NeverStop)
	return factory
}

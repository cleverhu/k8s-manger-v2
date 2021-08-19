package configs

import (
	"k8s-manger-v2/services"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var K8sClient *kubernetes.Clientset

type K8sConfig struct {
	DepHandler       *services.DepHandler       `inject:"-"`
	PodHandler       *services.PodHandler       `inject:"-"`
	RSHandler        *services.RSHandler        `inject:"-"`
	NSHandler        *services.NSHandler        `inject:"-"`
	EventHandler     *services.EventHandler     `inject:"-"`
	IngressHandler   *services.IngressHandler   `inject:"-"`
	ServiceHandler   *services.ServiceHandler   `inject:"-"`
	SecretHandler    *services.SecretHandler    `inject:"-"`
	ConfigMapHandler *services.ConfigMapHandler `inject:"-"`
}

func init() {
	K8sClient = initK8sClient()
}

func initK8sClient() *kubernetes.Clientset {
	cfg, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		log.Fatal(err)
	}
	cfg.Insecure = true
	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (this *K8sConfig) K8sClient() *kubernetes.Clientset {
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

	eventInformer := factory.Core().V1().Events().Informer()
	eventInformer.AddEventHandler(this.EventHandler)

	ingressInformer := factory.Networking().V1beta1().Ingresses().Informer()
	ingressInformer.AddEventHandler(this.IngressHandler)

	serviceInformer := factory.Core().V1().Services().Informer()
	serviceInformer.AddEventHandler(this.ServiceHandler)

	secretsInformer := factory.Core().V1().Secrets().Informer()
	secretsInformer.AddEventHandler(this.SecretHandler)

	configMapsInformer := factory.Core().V1().ConfigMaps().Informer()
	configMapsInformer.AddEventHandler(this.ConfigMapHandler)

	factory.Start(wait.NeverStop)
	return factory
}

package services

import (
	"k8s-manger-v2/wscore"
	corev1 "k8s.io/api/core/v1"
	"log"
)

type SecretHandler struct {
	SecretMap     *SecretMap     `inject:"-"`
	SecretService *SecretService `inject:"-"`
}

func (this *SecretHandler) OnAdd(obj interface{}) {
	secret := obj.(*corev1.Secret)
	this.SecretMap.Add(secret)

	wscore.ClientMap.SendAll("Secrets", secret.Namespace, this.SecretService.GetSecretsListByNS(secret.Namespace))
}
func (this *SecretHandler) OnUpdate(oldObj interface{}, newObj interface{}) {
	secret := newObj.(*corev1.Secret)
	err := this.SecretMap.Update(secret)
	if err != nil {
		log.Println(err)
	} else {
		wscore.ClientMap.SendAll("Secrets", secret.Namespace, this.SecretService.GetSecretsListByNS(secret.Namespace))
	}
}
func (this *SecretHandler) OnDelete(obj interface{}) {
	secret := obj.(*corev1.Secret)
	this.SecretMap.Delete(secret)
	wscore.ClientMap.SendAll("Secrets", secret.Namespace, this.SecretService.GetSecretsListByNS(secret.Namespace))
}

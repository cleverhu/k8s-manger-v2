package services

import (
	"context"
	"k8s-manger-v2/models"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapService struct {
	Client        *kubernetes.Clientset `inject:"-"`
	CommonService *CommonService        `inject:"-"`
	ConfigMapMap  *ConfigMapMap         `inject:"-"`
}

func NewConfigMapService() *ConfigMapService {
	return &ConfigMapService{}
}

func (this *ConfigMapService) GetConfigMapsListByNS(ns string) []*models.ConfigMap {

	secrets := this.ConfigMapMap.ListByNS(ns)
	ret := make([]*models.ConfigMap, 0)
	for _, secret := range secrets {
		ret = append(ret, &models.ConfigMap{
			Name:       secret.Name,
			NameSpace:  secret.Namespace,
			CreateTime: this.CommonService.TimeFormat(secret.CreationTimestamp.Time),
		})
	}
	return ret
}

func (this *ConfigMapService) PostConfigMap(postModel *models.PostConfigMapModel) error {
	cm := &corev1.ConfigMap{
		ObjectMeta: v1.ObjectMeta{
			Name:      postModel.Name,
			Namespace: postModel.NameSpace},
		Data: postModel.Data,
	}
	var err error
	if postModel.IsUpdate {
		_, err = this.Client.CoreV1().ConfigMaps(postModel.NameSpace).Update(context.Background(), cm, v1.UpdateOptions{})

	} else {
		_, err = this.Client.CoreV1().ConfigMaps(postModel.NameSpace).Create(context.Background(), cm, v1.CreateOptions{})
	}
	return err
}

func (this *ConfigMapService) GetConfigMap(ns, name string) *models.ConfigMap {

	cms := this.ConfigMapMap.ListByNS(ns)

	for _, cm := range cms {
		if cm.Name == name {
			return &models.ConfigMap{
				Name:       cm.Name,
				NameSpace:  cm.Namespace,
				CreateTime: this.CommonService.TimeFormat(cm.CreationTimestamp.Time),
				Data:       cm.Data,
			}
		}
	}
	return nil
}

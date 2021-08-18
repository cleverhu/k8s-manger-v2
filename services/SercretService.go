package services

import (
	"context"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/helpers"
	"k8s-manger-v2/models"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretService struct {
	Client        *kubernetes.Clientset `inject:"-"`
	CommonService *CommonService        `inject:"-"`
	SecretMap     *SecretMap            `inject:"-"`
}

func NewSecretService() *SecretService {
	return &SecretService{}
}

func (this *SecretService) GetSecretsListByNS(ns string) []*models.Secret {

	secrets, err := this.SecretMap.ListByNS(ns)
	goft.Error(err)
	ret := make([]*models.Secret, 0)
	for _, secret := range secrets {
		ret = append(ret, &models.Secret{
			Name:       secret.Name,
			NameSpace:  secret.Namespace,
			CreateTime: this.CommonService.TimeFormat(secret.CreationTimestamp.Time),
			Type:       models.SECRET_TYPE[string(secret.Type)],
		})
	}
	return ret
}

func (this *SecretService) PostSecret(postModel *models.PostSecretModel) error {
	secret := &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name:      postModel.Name,
			Namespace: postModel.NameSpace},
		Type:       corev1.SecretType(postModel.Type),
		StringData: postModel.Data,
	}
	_, err := this.Client.CoreV1().Secrets(postModel.NameSpace).Create(context.Background(), secret, v1.CreateOptions{})
	return err
}

func (this *SecretService) GetSecret(ns, name string) *models.Secret {

	secrets, err := this.SecretMap.ListByNS(ns)
	goft.Error(err)

	for _, secret := range secrets {
		if secret.Name == name {
			return &models.Secret{
				Name:       secret.Name,
				NameSpace:  secret.Namespace,
				CreateTime: this.CommonService.TimeFormat(secret.CreationTimestamp.Time),
				Type:       models.SECRET_TYPE[string(secret.Type)],
				Data:       secret.Data,
				ExtData:    this.ParseIfTLS(string(secret.Type), secret.Data),
			}
		}
	}
	return nil
}

//解析 （如类型是 tls 的secret)
func (this *SecretService) ParseIfTLS(t string, data map[string][]byte) interface{} {
	if t == "kubernetes.io/tls" {
		if crt, ok := data["tls.crt"]; ok {
			crtModel := helpers.ParseCert(crt)
			if crtModel != nil {
				return crtModel
			}
		}
	}
	return struct{}{}

}

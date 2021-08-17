package services

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/models"
)

type IngressService struct {
	CommonService *CommonService `inject:"-"`
	IngressMap    *IngressMap    `inject:"-"`
}

func NewIngressService() *IngressService {
	return &IngressService{}
}

func (this *IngressService) GetIngressListByNS(ns string) []*models.Ingress {
	ingresses, err := this.IngressMap.ListByNS(ns)
	goft.Error(err)
	ret := make([]*models.Ingress, 0)
	for _, ingress := range ingresses {
		ret = append(ret, &models.Ingress{
			Name:       ingress.Name,
			NameSpace:  ingress.Namespace,
			CreateTime: this.CommonService.TimeFormat(ingress.CreationTimestamp.Time),
		})
	}
	return ret
}

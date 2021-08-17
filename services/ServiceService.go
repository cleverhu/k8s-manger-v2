package services

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/models"
)

type ServiceService struct {
	CommonService *CommonService `inject:"-"`
	ServiceMap    *ServiceMap    `inject:"-"`
}

func NewServiceService() *ServiceService {
	return &ServiceService{}
}

func (this *ServiceService) GetServiceListByNS(ns string) []*models.Service {
	services, err := this.ServiceMap.ListByNS(ns)
	goft.Error(err)
	ret := make([]*models.Service, 0)
	for _, svc := range services {
		ret = append(ret, &models.Service{
			Name:       svc.Name,
			NameSpace:  svc.Namespace,
			CreateTime: this.CommonService.TimeFormat(svc.CreationTimestamp.Time),
		})
	}
	return ret
}

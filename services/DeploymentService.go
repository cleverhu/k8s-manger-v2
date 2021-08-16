package services

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/core"
	"k8s-manger-v2/models"
)

type DeploymentService struct {
	CommonSvc *CommonService      `inject:"-"`
	DepMap    *core.DeploymentMap `inject:"-"`
}

func NewDeploymentService() *DeploymentService {
	return &DeploymentService{}
}

func (this *DeploymentService) ListAll(namespace string) []*models.Deployment {
	ret := make([]*models.Deployment, 0)
	deps, err := this.DepMap.ListByNS(namespace)
	goft.Error(err)
	for _, dep := range deps {
		tmp := &models.Deployment{
			NameSpace: dep.Namespace,
			Name:      dep.Name,
			Replicas:  [3]int32{dep.Status.Replicas, dep.Status.AvailableReplicas, dep.Status.UnavailableReplicas},
			Images:    this.CommonSvc.GetImages(*dep),
		}
		ret = append(ret, tmp)
	}
	return ret
}

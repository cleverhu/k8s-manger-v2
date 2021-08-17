package services

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/models"
	v1 "k8s.io/api/apps/v1"
)

type DeploymentService struct {
	CommonService *CommonService `inject:"-"`
	DepMap        *DeploymentMap `inject:"-"`
	PodService    *PodService    `inject:"-"`
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
			NameSpace:  dep.Namespace,
			Name:       dep.Name,
			Replicas:   [3]int32{dep.Status.Replicas, dep.Status.AvailableReplicas, dep.Status.UnavailableReplicas},
			Images:     this.CommonService.GetImagesByDep(*dep),
			CreateTime: this.CommonService.TimeFormat(dep.CreationTimestamp.Time),
			//Pods: this.PodService.GetPodsByDep(*dep),
			IsComplete: this.getDeploymentIsComplete(dep),
			Message:    this.getDeploymentCondition(dep),
		}
		ret = append(ret, tmp)
	}
	return ret
}

func (this *DeploymentService) Detail(namespace string, name string) *models.Deployment {
	deploy, err := this.DepMap.Get(namespace, name)
	goft.Error(err)
	return &models.Deployment{
		Name:       name,
		NameSpace:  namespace,
		Images:     this.CommonService.GetImagesByDep(*deploy),
		CreateTime: this.CommonService.TimeFormat(deploy.CreationTimestamp.Time),
		Pods:       this.PodService.GetPodsByDep(*deploy),
		Replicas:   [3]int32{deploy.Status.Replicas, deploy.Status.AvailableReplicas, deploy.Status.UnavailableReplicas},
	}
}

func (*DeploymentService) getDeploymentCondition(dep *v1.Deployment) string {
	for _, item := range dep.Status.Conditions {
		if string(item.Type) == "Available" && string(item.Status) != "True" {
			return item.Message
		}
	}
	return ""
}

func (*DeploymentService) getDeploymentIsComplete(dep *v1.Deployment) bool {
	return dep.Status.Replicas == dep.Status.AvailableReplicas
}

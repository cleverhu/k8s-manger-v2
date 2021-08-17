package services

import (
	"github.com/shenyisyn/goft-gin/goft"
	"k8s-manger-v2/models"
	v1 "k8s.io/api/apps/v1"
)

type PodService struct {
	PodMap        *PodMap        `inject:"-"`
	CommonService *CommonService `inject:"-"`
	RSMap         *RSMap         `inject:"-"`
	EventMap      *EventMap      `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}

func (this *PodService) GetPodsByDep(dep v1.Deployment) []*models.Pod {
	rsLabelsMap, err := this.RSMap.GetRsLabelsByDeployment(&dep)
	goft.Error(err)
	pods, err := this.PodMap.ListByRsLabelsAndNS(dep.Namespace, rsLabelsMap)
	goft.Error(err)
	ret := make([]*models.Pod, 0)
	for _, pod := range pods {
		ret = append(ret, &models.Pod{
			Name:       pod.Name,
			NameSpace:  pod.Namespace,
			Images:     this.CommonService.GetImagesByPod(pod.Spec.Containers),
			NodeName:   pod.Spec.NodeName,
			CreateTime: this.CommonService.TimeFormat(pod.CreationTimestamp.Time),
			IPs:        []string{pod.Status.PodIP, pod.Status.HostIP},
		})
	}
	return ret
}

func (this *PodService) GetPodsListByNS(ns string) []*models.Pod {

	pods, err := this.PodMap.ListByNS(ns)
	goft.Error(err)
	ret := make([]*models.Pod, 0)
	for _, pod := range pods {
		ret = append(ret, &models.Pod{
			Name:       pod.Name,
			NameSpace:  pod.Namespace,
			Images:     this.CommonService.GetImagesByPod(pod.Spec.Containers),
			NodeName:   pod.Spec.NodeName,
			CreateTime: this.CommonService.TimeFormat(pod.CreationTimestamp.Time),
			IPs:        []string{pod.Status.PodIP, pod.Status.HostIP},
			IsReady:    this.CommonService.PosIsReady(pod),
			Message:    this.EventMap.GetMessage(pod.Namespace, "Pod", pod.Name),
		})
	}
	return ret
}

package services

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"time"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (this *CommonService) GetImagesByDep(dep v1.Deployment) string {
	return this.GetImagesByPod(dep.Spec.Template.Spec.Containers)
}

func (this *CommonService) GetImagesByPod(containers []corev1.Container) string {
	images := containers[0].Image
	if imgLen := len(containers); imgLen > 1 {
		images += fmt.Sprintf("+其他%d个镜像", imgLen-1)
	}
	return images
}

func (this *CommonService) TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (this *CommonService) IsValidLabel(m1, m2 map[string]string) bool {
	for key := range m2 {
		if m2[key] != m1[key] {
			return false
		}
	}

	return true
}

func (*CommonService) PosIsReady(pod *corev1.Pod) bool {
	if pod.Status.Phase != "Running" {
		return false
	}
	for _, condition := range pod.Status.Conditions {
		if condition.Status != "True" {
			return false
		}
	}
	for _, rg := range pod.Spec.ReadinessGates {
		for _, condition := range pod.Status.Conditions {
			if condition.Type == rg.ConditionType && condition.Status != "True" {
				return false
			}
		}
	}
	return true
}

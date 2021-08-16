package services

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (this *CommonService) GetImages(dep v1.Deployment) string {
	return this.getImages(dep.Spec.Template.Spec.Containers)
}

func (this *CommonService) getImages(containers []corev1.Container) string {
	images := ""
	if imgLen := len(containers); imgLen > 1 {
		images += fmt.Sprintf("+其他%d个镜像", imgLen-1)
	}
	return images
}

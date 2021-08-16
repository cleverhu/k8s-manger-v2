package configs

import "k8s-manger-v2/services"

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (*ServiceConfig) DeploymentService() *services.DeploymentService {
	return services.NewDeploymentService()
}

func (*ServiceConfig) CommonService() *services.CommonService {
	return services.NewCommonService()
}

func (*ServiceConfig) PodService() *services.PodService {
	return services.NewPodService()
}

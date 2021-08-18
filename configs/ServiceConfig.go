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

func (*ServiceConfig) NSService() *services.NSService {
	return services.NewNSService()
}

func (*ServiceConfig) IngressService() *services.IngressService {
	return services.NewIngressService()
}

func (*ServiceConfig) ServiceService() *services.ServiceService {
	return services.NewServiceService()
}

func (*ServiceConfig) SecretService() *services.SecretService {
	return services.NewSecretService()
}

func (*ServiceConfig) ConfigMapService() *services.ConfigMapService {
	return services.NewConfigMapService()
}

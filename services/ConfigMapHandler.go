package services

import (
	"k8s-manger-v2/wscore"
	corev1 "k8s.io/api/core/v1"
)

type ConfigMapHandler struct {
	ConfigMapMap     *ConfigMapMap     `inject:"-"`
	ConfigMapService *ConfigMapService `inject:"-"`
}

func (this *ConfigMapHandler) OnAdd(obj interface{}) {
	cm := obj.(*corev1.ConfigMap)
	this.ConfigMapMap.Add(cm)
	wscore.ClientMap.SendAll("ConfigMaps", cm.Namespace, this.ConfigMapService.GetConfigMapsListByNS(cm.Namespace))

}
func (this *ConfigMapHandler) OnUpdate(oldObj interface{}, newObj interface{}) {

	oldCM := oldObj.(*corev1.ConfigMap)
	newCm := newObj.(*corev1.ConfigMap)
	this.ConfigMapMap.Update(newCm)
	if this.DataIsEqual(oldCM, newCm) {
		wscore.ClientMap.SendAll("ConfigMaps", newCm.Namespace, this.ConfigMapService.GetConfigMapsListByNS(newCm.Namespace))
	}

}
func (this *ConfigMapHandler) OnDelete(obj interface{}) {
	cm := obj.(*corev1.ConfigMap)
	this.ConfigMapMap.Delete(cm)
	wscore.ClientMap.SendAll("ConfigMaps", cm.Namespace, this.ConfigMapService.GetConfigMapsListByNS(cm.Namespace))
}

func (this *ConfigMapHandler) DataIsEqual(cm1 *corev1.ConfigMap, cm2 *corev1.ConfigMap) bool {
	m1 := cm1.Data
	m2 := cm2.Data
	if len(m1) != len(m2) {
		return false
	}

	for key := range m1 {
		if m1[key] != m2[key] {
			return false
		}
	}

	return true
}

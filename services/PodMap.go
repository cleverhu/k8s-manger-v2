package services

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type PodMap struct {
	CommonService *CommonService `inject:"-"`
	data          sync.Map       //key ns value []*corev1.Pod
}

func (this *PodMap) Add(pod *corev1.Pod) {
	key := pod.Namespace
	if value, ok := this.data.Load(key); ok {
		value = append(value.([]*corev1.Pod), pod)
		this.data.Store(key, value)
	} else {
		this.data.Store(key, []*corev1.Pod{pod})
	}
}

func (this *PodMap) Delete(pod *corev1.Pod) {
	key := pod.Namespace
	if value, ok := this.data.Load(key); ok {
		for index, p := range value.([]*corev1.Pod) {
			if p.Name == pod.Name {
				value = append(value.([]*corev1.Pod)[0:index], value.([]*corev1.Pod)[index+1:]...)
				this.data.Store(key, value)
				return
			}
		}
	}
}

func (this *PodMap) Update(pod *corev1.Pod) error {
	key := pod.Namespace
	if value, ok := this.data.Load(key); ok {

		for index, p := range value.([]*corev1.Pod) {
			if p.Name == pod.Name {
				value.([]*corev1.Pod)[index] = pod
				this.data.Store(key, value)
				return nil
			}
		}
	}
	return fmt.Errorf("pod-%s not found", pod.Name)
}

func (this *PodMap) ListByNS(ns string) ([]*corev1.Pod, error) {

	if ns != "all-namespaces" {
		if list, ok := this.data.Load(ns); ok {
			return list.([]*corev1.Pod), nil
		}
	} else {
		ret := make([]*corev1.Pod, 0)
		this.data.Range(func(key, value interface{}) bool {
			for _, dep := range value.([]*corev1.Pod) {
				ret = append(ret, dep)
			}
			sort.Slice(ret, func(i, j int) bool {
				if ret[i].Namespace == ret[j].Namespace {
					return ret[i].Name < ret[j].Name
				} else {
					return ret[i].Namespace < ret[j].Namespace
				}
			})
			return true
		})
		return ret, nil
	}

	return nil, nil
	//return nil, fmt.Errorf("pods not found")
}

func (this *PodMap) ListByRsLabelsAndNS(ns string, rsLabels []map[string]string) ([]*corev1.Pod, error) {
	pods, err := this.ListByNS(ns)
	if err != nil {
		return nil, err
	}
	ret := make([]*corev1.Pod, 0)
	for _, pod := range pods {
		for _, rLabel := range rsLabels {
			if this.CommonService.IsValidLabel(pod.Labels, rLabel) {
				ret = append(ret, pod)
			}
		}
	}
	return ret, nil
}

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/utils/ptr"
)

var (
	_ EnhancedCRD = (*CsNodeDaemon)(nil)
)

func (cr *CsNodeDaemon) AsOwner() []metav1.OwnerReference {
	return []metav1.OwnerReference{
		{
			APIVersion:         cr.APIVersion,
			Kind:               cr.Kind,
			Name:               cr.Name,
			UID:                cr.UID,
			Controller:         ptr.To(true),
			BlockOwnerDeletion: ptr.To(true),
		},
	}
}

func (cr *CsNodeDaemon) PodLabels() map[string]string {
	return labels.Merge(CommonLabels(), map[string]string{
		"app.kubernetes.io/name":      "node-daemon",
		"app.kubernetes.io/part-of":   "node-daemon",
		"app.kubernetes.io/instance":  cr.Name,
		"app.kubernetes.io/component": "node-daemon",
	})
}

func (cr *CsNodeDaemon) AllLabels() map[string]string {
	return labels.Merge(CommonLabels(), map[string]string{
		"app.kubernetes.io/name":      "node-daemon",
		"app.kubernetes.io/part-of":   "node-daemon",
		"app.kubernetes.io/instance":  cr.Name,
		"app.kubernetes.io/component": "node-daemon",
	})
}

func (cr *CsNodeDaemon) SelectorLabels() map[string]string {
	return labels.Merge(CommonLabels(), map[string]string{
		"app.kubernetes.io/name":      "node-daemon",
		"app.kubernetes.io/part-of":   "node-daemon",
		"app.kubernetes.io/instance":  cr.Name,
		"app.kubernetes.io/component": "node-daemon",
	})
}

func (cr *CsNodeDaemon) GetServiceAccountName() string {
	return "default"
}

func (cr *CsNodeDaemon) GetNamespaceName() string {
	return ""
}

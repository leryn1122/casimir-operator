package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type EnhancedCRD interface {
	AsOwner() []metav1.OwnerReference
	PodLabels() map[string]string
	AllLabels() map[string]string
	SelectorLabels() map[string]string
	GetServiceAccountName() string
	GetNamespaceName() string
}

func CommonLabels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/version":    "0.1.0",
		"app.kubernetes.io/managed-by": "casimir-operator",
	}
}

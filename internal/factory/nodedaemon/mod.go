package nodedaemon

import (
	"context"
	"fmt"
	csv1alpha1 "github.com/leryn1122/casimir-operator/v2/api/v1alpha1"
	"github.com/leryn1122/casimir-operator/v2/internal/factory/reconcile"
	"github.com/leryn1122/casimir-operator/v2/schema/config"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateOrUpdateCsNodeDaemon(ctx context.Context, cr *csv1alpha1.CsNodeDaemon, client client.Client, config *config.OperatorConfig) error {
	_, err := newDaemonSetForCsNodeDaemon(cr, config)
	if err != nil {

	}

	newDaemonSet, err := newDaemonSetForCsNodeDaemon(cr, config)
	if err != nil {
		return fmt.Errorf("failed to generate CsNodeDaemon daemonSet with name %s: %w", cr.Name, err)
	}

	options := reconcile.DaemonSetOptions{}

	err = reconcile.HandleDaemonSetUpdate(ctx, client, options, newDaemonSet, config)
	if err != nil {
		return fmt.Errorf("failed to reconcile CsNodeDaemon daemonSet with name %s: %w", cr.Name, err)
	}
	return nil
}

func newDaemonSetForCsNodeDaemon(cr *csv1alpha1.CsNodeDaemon, config *config.OperatorConfig) (*appsv1.DaemonSet, error) {
	podSpec := v1.PodSpec{
		NodeSelector: map[string]string{
			"kubernetes.io/os": "linux",
		},
		Containers: []v1.Container{
			{
				Name: "node-daemon",
				VolumeMounts: []v1.VolumeMount{
					{
						Name:      "host-localtime",
						MountPath: "/etc/localtime",
						ReadOnly:  true,
					},
				},
			},
		},
		Volumes: []v1.Volume{
			{
				Name: "host-localtime",
				VolumeSource: v1.VolumeSource{
					HostPath: &v1.HostPathVolumeSource{
						Path: "/etc/localtime",
					},
				},
			},
		},
		Tolerations: []v1.Toleration{
			{
				Operator: v1.TolerationOpExists,
				Effect:   v1.TaintEffectNoSchedule,
			},
			{
				Operator: v1.TolerationOpExists,
				Effect:   v1.TaintEffectNoExecute,
			},
		},
	}

	daemonSet := appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:            cr.Name,
			Namespace:       config.ManagedNamespace,
			Labels:          cr.AllLabels(),
			Annotations:     map[string]string{},
			OwnerReferences: cr.AsOwner(),
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/name":       "node-daemon",
					"app.kubernetes.io/part-of":    "node-daemon",
					"app.kubernetes.io/component":  "node-daemon",
					"app.kubernetes.io/version":    "0.1.0",
					"app.kubernetes.io/managed-by": "casimir-operator",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: cr.PodLabels(),
				},
				Spec: podSpec,
			},
		},
	}

	return &daemonSet, nil
}

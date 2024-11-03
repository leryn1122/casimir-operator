package reconcile

import (
	"context"
	"github.com/leryn1122/casimir-operator/v2/schema/config"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DaemonSetOptions struct {
	SelectorLabels func() map[string]string
}

func waitForDaemonSetReady(ctx context.Context, client client.Client, daemonSet *appsv1.DaemonSet, config *config.OperatorConfig) error {
	return nil
}

func HandleDaemonSetUpdate(ctx context.Context, client client.Client, options DaemonSetOptions, newDaemonSet *appsv1.DaemonSet, config *config.OperatorConfig) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		var currentDaemonSet appsv1.DaemonSet
		if err := client.Get(ctx, types.NamespacedName{
			Name:      newDaemonSet.Name,
			Namespace: newDaemonSet.Namespace,
		}, &currentDaemonSet); err != nil {
			if errors.IsNotFound(err) {
				if err = client.Create(ctx, newDaemonSet); err != nil {
					return err
				}
			}
			return waitForDaemonSetReady(ctx, client, newDaemonSet, config)
		}
		return nil
	})
}

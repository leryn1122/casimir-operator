package manager

import (
	"context"
	"github.com/leryn1122/casimir-operator/v2/pkg/controllers"
	"github.com/leryn1122/casimir-operator/v2/pkg/logger"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/flowcontrol"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
)

func NewManager(ctx context.Context) error {

	logger.InitializeLogger()

	config := ctrl.GetConfigOrDie()
	config.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(
		float32(32),
		100,
	)

	manager, err := ctrl.NewManager(config, ctrl.Options{})
	if err != nil {
		ctrl.Log.Error(err, "failed to create manager")
	}

	// Add `/readyz` and `healthz` check
	if err = manager.AddReadyzCheck("readyz", newReadyHandler(ctx)); err != nil {
		return errors.Wrap(err, "failed to add readyz check")
	}
	if err = manager.AddHealthzCheck("healthz", newHealthHandler(ctx)); err != nil {
		return errors.Wrap(err, "failed to add health check")
	}

	k8sClient, err := kubernetes.NewForConfig(manager.GetConfig())
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes client")
	}

	k8sVersion, err := k8sClient.ServerVersion()
	if err != nil {
		return errors.Wrap(err, "failed to get kubernetes server version")
	}
	ctrl.Log.Info("kubernetes version: %s", k8sVersion.String())

	watch, err := client.NewWithWatch(manager.GetConfig(), client.Options{
		Scheme: runtime.NewScheme(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to create watch client")
	}
	converter, err := controllers.NewConvertController(ctx, k8sClient, watch)
	if err != nil {
		return errors.Wrap(err, "failed to create convert controller")
	}

	if err = manager.Add(converter); err != nil {
		return errors.Wrap(err, "failed to add controller")
	}

	if err = manager.Start(ctx); err != nil {
		return errors.Wrap(err, "failed to start manager")
	}

	return nil
}

func addWebhookConfigurations(manager ctrl.Manager) error {
	f := func(objects []client.Object) error {
		var err error
		for _, object := range objects {
			if err = ctrl.NewWebhookManagedBy(manager).For(object).Complete(); err != nil {
				return err
			}
		}
		return nil
	}

	return f([]client.Object{})
}

func newReadyHandler(ctx context.Context) healthz.Checker {
	return func(request *http.Request) error {
		return nil
	}
}

func newHealthHandler(ctx context.Context) healthz.Checker {
	return func(request *http.Request) error {
		return nil
	}
}

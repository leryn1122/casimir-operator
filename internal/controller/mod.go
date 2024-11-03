package controller

import (
	"context"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sync"
	"time"
)

var (
	cacheSyncTimeout = ptr.To(3 * time.Minute)
	maxConcurrency   = ptr.To(5)

	optionsInit    sync.Once
	defaultOptions *controller.Options
)

func getDefaultControllerOptions() controller.Options {
	optionsInit.Do(func() {
		defaultOptions = &controller.Options{
			RateLimiter:             workqueue.NewTypedItemExponentialFailureRateLimiter[reconcile.Request](2*time.Second, 2*time.Minute),
			CacheSyncTimeout:        *cacheSyncTimeout,
			MaxConcurrentReconciles: *maxConcurrency,
		}
	})
	return *defaultOptions
}

func handleReconcileError(ctx context.Context, client client.Client) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

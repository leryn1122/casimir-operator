/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"github.com/leryn1122/casimir-operator/v2/internal/factory/nodedaemon"
	"github.com/leryn1122/casimir-operator/v2/schema/config"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"

	operatorv1alpha1 "github.com/leryn1122/casimir-operator/v2/api/v1alpha1"
)

var _ reconcile.Reconciler = (*CsNodeDaemonReconciler)(nil)

// CsNodeDaemonReconciler reconciles a CsNodeDaemon object
type CsNodeDaemonReconciler struct {
	client.Client
	OriginScheme *runtime.Scheme
	Config       *config.OperatorConfig
}

// +kubebuilder:rbac:groups=operator.leryn.io,resources=csnodedaemons,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.leryn.io,resources=csnodedaemons/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=operator.leryn.io,resources=csnodedaemons/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CsNodeDaemon object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *CsNodeDaemonReconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx).WithName("CsNodeDaemonReconciler")
	log.Info("Reconciling CsNodeDaemon")

	instance := &operatorv1alpha1.CsNodeDaemon{}

	defer func() {

	}()

	if err := r.Get(ctx, types.NamespacedName{
		Name:      request.Name,
		Namespace: r.Config.ManagedNamespace,
	}, instance); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	r.Client.Scheme().Default(instance)

	if err := nodedaemon.CreateOrUpdateCsNodeDaemon(ctx, instance, r, r.Config); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CsNodeDaemonReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		Named("CsNodeDaemon").
		For(&operatorv1alpha1.CsNodeDaemon{}).
		Owns(&v1.DaemonSet{}).
		WithOptions(getDefaultControllerOptions()).
		Complete(r)
}

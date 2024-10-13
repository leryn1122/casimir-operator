package main

import (
	"context"
	"os"

	"github.com/leryn1122/casimir-operator/v2/pkg/manager"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func main() {
	ctrl.Log.Info("starting controller")

	ctx, cancel := context.WithCancel(context.Background())
	stop := signals.SetupSignalHandler()
	go func() {
		<-stop.Done()
		cancel()
	}()

	err := manager.NewManager(ctx)
	if err != nil {
		ctrl.Log.Error(err, "failed to start controller")
		os.Exit(1)
	}

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))
	ctrl.Log.Info("stop manager")
}

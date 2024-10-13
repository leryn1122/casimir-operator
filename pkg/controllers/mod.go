package controllers

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ manager.Runnable = (*ConverterController)(nil)

type ConverterController struct {
	ctx       context.Context
	k8sClient *kubernetes.Clientset
	watch     client.WithWatch
}

func NewConvertController(ctx context.Context, k8sClient *kubernetes.Clientset, watch client.WithWatch) (*ConverterController, error) {
	c := &ConverterController{
		ctx:       ctx,
		k8sClient: k8sClient,
		watch:     watch,
	}
	return c, nil
}

func (c *ConverterController) Start(context.Context) error {
	return nil
}

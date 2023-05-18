package objects

import (
	"context"

	"k8s.io/client-go/kubernetes"
)


func NewObject(ctx context.Context, clientset *kubernetes.Clientset) (*Object) {
	return &Object{
		ctx: ctx,
		clientset: clientset,
	}
}
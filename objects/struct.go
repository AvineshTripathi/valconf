package objects

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

type Object struct {
	ctx context.Context
	clientset *kubernetes.Clientset
}
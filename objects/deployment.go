package objects

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func(object *Object) Deployment(deployList *appsv1.DeploymentList) (deployments []map[string]interface{}) {
	
	for _, item := range deployList.Items {

		deployConfig, err := object.clientset.AppsV1().Deployments(item.Namespace).Get(object.ctx, item.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
		}

		deployment := map[string]interface{} {
			"name": deployConfig.Name,
			"replica": *deployConfig.Spec.Replicas,
			"strategy": deployConfig.Spec.Strategy.Type,
			"labels": deployConfig.Labels,
		}

		deployments = append(deployments, deployment)
	}

	return 
}
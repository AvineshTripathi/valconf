package objects

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func(object *Object) Service(serviceList *corev1.ServiceList) (Svcs []map[string]interface{}){
	for _, item := range serviceList.Items {

		deployConfig, err := object.clientset.CoreV1().Services(item.Namespace).Get(object.ctx, item.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
		}

		ports := GetPorts(deployConfig.Spec.Ports)

		svc := map[string]interface{}{
			"name": item.Name,
			"ports": ports,
		}
		
		Svcs = append(Svcs, svc)
	}

	return 

}


func GetPorts(p []corev1.ServicePort) (arr []int32) {
	for _, item := range p {
		arr = append(arr, item.Port)
	}
	return 
}
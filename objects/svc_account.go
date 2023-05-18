package objects

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func(object *Object) ServiceAccount(svcAccountList *corev1.ServiceAccountList) (SvcAccs []map[string]interface{}) {
	for _, item := range svcAccountList.Items {
		
		svcAccConfig, err := object.clientset.CoreV1().ServiceAccounts(item.Namespace).Get(object.ctx, item.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
		}
		
		svcAcc := map[string]interface{}{
			"name": item.Name,
			"namespace": svcAccConfig.Namespace,
		}

		SvcAccs = append(SvcAccs, svcAcc)
	}

	return
}

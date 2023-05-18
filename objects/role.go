package objects

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/rbac/v1"
)

func(object *Object) Role(roleList *v1.RoleList) (Roles []map[string]interface{}) {
	for _, item := range roleList.Items {

		rolesConfig, err := object.clientset.RbacV1().Roles(item.Namespace).Get(object.ctx, item.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
		}
		var a []map[string]interface{}
		for _, i := range rolesConfig.Rules {
			a = append(a, map[string]interface{}{
				"resources": i.Resources,
				"verbs": i.Verbs,
			})
		}
		
		Roles = append(Roles, map[string]interface{}{
			"name": item.Name,
			"rules": a,
		})
	}

	return  
}
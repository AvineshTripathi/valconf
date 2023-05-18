package objects

import (
	//"bytes"
	//"bytes"

	//"encoding/json"

	//"encoding/json"
	//"strings"

	//"encoding/json"
	//"io/ioutil"

	//"encoding/json"
	"fmt"
	//"os"

	//"gopkg.in/yaml.v3"
	//`:wqa`"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

func(object *Object) Pod(podsList *corev1.PodList) (pods []map[string]interface{}) {
	for _, item := range podsList.Items {
		namespace := item.Namespace
		name := item.Name

		podConfig, err := object.clientset.CoreV1().Pods(namespace).Get(object.ctx, name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
		}


		// data, err := json.Marshal(podConfig)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// var input interface{}
		// d := json.NewDecoder(bytes.NewBufferString(string(data)))
		// // Numeric values must be represented using json.Number.
		// d.UseNumber()
		// if err := d.Decode(&input); err != nil {
		// 	fmt.Println(err)
		// }

		c := Containers(podConfig)
		input := map[string]interface{}{
			"pod": c,
			//"volumes": Volumes(podConfig), map->pod->privi | 
		}
		pods = append(pods, input)
		// policyKey := fmt.Sprintf("%s.rego", "eval")
		// compiler, err := ast.CompileModules(map[string]string{
		// 	policyKey: utils.RegoUtils,
		// })
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// pkgName := compiler.Modules[policyKey].Package.Path.String()

		// regoFunc := make([]func(r *rego.Rego), 0)

		// regoFunc = append(regoFunc, rego.Query(fmt.Sprintf("%s", pkgName)))

		// regoFunc = append(regoFunc, rego.Compiler(compiler))
		// regoFunc = append(regoFunc, rego.Input(input))

		// rego := rego.New(regoFunc...)

		// res, err := rego.Eval(object.ctx)

		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(res)
		// fmt.Println(res[0].Expressions[0])
		// fmt.Println(res[0].Bindings)

		// c := PodConfig{
		// 	APIVersion: podConfig.APIVersion,
		// 	Kind:       podConfig.Kind,
		// 	Metadata:   podConfig.ObjectMeta,
		// 	Spec:       podConfig.Spec,
		// }

		//fmt.Println(c.Metadata.Annotations)
		// r := rego.New(
		// 	rego.Load([]string{"policy.rego"}, nil),
		// 	rego.Query("input.allow"),
		// )

		// fmt.Println(r)

		// prep, err := r.PrepareForEval(ctx)
		// if err != nil {
		// 	fmt.Println(err, "prep")
		// }

		// rs, err := prep.Eval(ctx, rego.EvalInput(c))
		// if err != nil {
		// 	fmt.Println(err, " eval")
		// }

		// fmt.Println(rs)
		
	}
	return 
}


func Containers(podConfig *corev1.Pod) (containers []map[string]interface{}) {
	
	for _, item := range podConfig.Spec.Containers {
		
		var resP string
		var resL = "false"

		if item.SecurityContext == nil || item.SecurityContext.Privileged == nil {
			resP = "false"
		} else if(*item.SecurityContext.Privileged){
			resP = "true"
		}
		
		if item.Resources.Limits != nil {
			resL = "true"
		}

		a := map[string]interface{}{
			"name": podConfig.Name,
			"privileged": resP,
			"limits": resL,
		}

		containers = append(containers, a)
	}

	return 
}



// Not working! Debugging required
func Volumes(podConfig *corev1.Pod) (volumes []map[string]interface{}) {
		
		for _, vol := range podConfig.Spec.Volumes {
	
			a := map[string]interface{}{
				"name": vol.Name,
				"path": vol.HostPath.Path,
			}

			volumes = append(volumes, a)
		}

		return 
}

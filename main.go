package main

import (
	"context"
	"fmt"
	"os"

	"github.com/AvineshTripathi/valconf/objects"
	"github.com/AvineshTripathi/valconf/validate"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


func main() {

	ctx := context.Background()

	kubeconfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		fmt.Println(err)
	}

	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Println(err)
	}

	object := objects.NewObject(ctx, clientset)
	
	podsList, err := clientset.CoreV1().Pods("").List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	deployList, err := clientset.AppsV1().Deployments("").List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	serviceList, err := clientset.CoreV1().Services("").List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	svcAccountList, err := clientset.CoreV1().ServiceAccounts("").List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	roleList, err := clientset.RbacV1().Roles("").List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// optimised
	// channels := [5]chan map[string]interface{}{}

	// for i := 0; i<5; i++ {
	// 	channels[i] = make(chan map[string]interface{}, 1)
	// 	go func(index int) {
	// 		var result map[string]interface{}
	// 		switch index {
	// 		case 0:
	// 			result = map[string]interface{}{"pod": object.Pod(podsList)}
	// 		case 1:
	// 			result = map[string]interface{}{"deploy": object.Deployment(deployList)}
	// 		case 2:
	// 			result = map[string]interface{}{"svc": object.Service(serviceList)}
	// 		case 3:
	// 			result = map[string]interface{}{"svc_account": object.ServiceAccount(svcAccountList)}
	// 		case 4:
	// 			result = map[string]interface{}{"role": object.Role(roleList)}
	// 		}
	// 		channels[index] <- result
	// 	}(i)
	// }

	// data := make([]map[string]interface{}, 5)
	// for i, _ := range channels {
	// 	result := <- channels[i]
	// 	fmt.Println(result)
	// 	data = append(data, result)
	// }
	
	podConf := object.Pod(podsList)
	deployConf := object.Deployment(deployList)
	svcConf := object.Service(serviceList)
	svcAccConf := object.ServiceAccount(svcAccountList)
	roleConf := object.Role(roleList)

 	input := map[string]interface{}{
		"pod": podConf,
		"deploy": deployConf,
		"svc": svcConf,
		"svc_account": svcAccConf,
		"role": roleConf,
	}

	// fmt.Println(len(input))
	// fmt.Println(input["pod"])
	// fmt.Println()
	// fmt.Println(input["deploy"])
	// fmt.Println()
	// fmt.Println(input["svc"])
	// fmt.Println()
	// fmt.Println(input["svc_account"])
	// fmt.Println()
	// fmt.Println(input["role"])

	validate.Rego(ctx, input)
	
	//fmt.Println(input["pod"], object.Pod(podsList))
}



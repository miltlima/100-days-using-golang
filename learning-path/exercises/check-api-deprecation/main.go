package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// create a Kubernetes client using in-cluster configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// specify the API resource you want to check
	resource := "deployments"

	// retrieve the resource object and check if it's deprecated
	apiResourceList, err := clientset.Discovery().ServerResourcesForGroupVersion("apps/v1")
	if err != nil {
		panic(err.Error())
	}
	for _, apiResource := range apiResourceList.APIResources {
		if apiResource.Name == resource {
			if apiResource.Deprecated != nil && *apiResource.Deprecated {
				fmt.Printf("%s is deprecated\n", resource)
			} else {
				fmt.Printf("%s is not deprecated\n", resource)
			}
			break
		}
	}
}

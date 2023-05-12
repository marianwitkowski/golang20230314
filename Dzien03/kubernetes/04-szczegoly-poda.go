package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeConfig() (*rest.Config, error) {
	kubeconfig := "/Users/marian/.kube/config" //os.Getenv("KUBECONFIG")
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return nil, fmt.Errorf("KUBECONFIG nie ustawiony")
}

func main() {

	config, err := getKubeConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	namespace := "default"
	podName := "ngix-szkolenie"

	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(pod.ObjectMeta.UID)
	fmt.Println(pod.Status.Phase)
	for _, container := range pod.Spec.Containers {
		fmt.Println(container.Name, container.Image)
		for _, port := range container.Ports {
			fmt.Printf("Port nr %d\n", port.ContainerPort)
		}
	}

}

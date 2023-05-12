package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
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

	nginxPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ngix-szkolenie",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx-szkol-1",
					Image: "nginx:latest",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}

	pod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), nginxPod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pod %s utworzony\n", pod.Name)

}

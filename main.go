package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

type Configuration struct {
	kubeConfigPath string
	format         string
}

func getenv(key string, defaultValue string) (value string, err error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	if len(defaultValue) == 0 {
		return "", fmt.Errorf("%s is not set and does not have a default value", key)
	}
	return defaultValue, nil
}

func configure() (config Configuration) {
	kubeConfigPath, err := getenv("KUBECONFIG", "")
	if err != nil {
		log.Fatal(err)
	}
	config.kubeConfigPath = kubeConfigPath
	config.format, _ = getenv("OUTPUT_FORMAT", "json")
	return
}

func main() {
	config := configure()
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", config.kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("Pod name=/%s\n", pod.GetName())
	}
}

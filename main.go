package main

import (
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	kubeConfigPath string
	format         string
	isInCluster    bool
}

type Output struct {
	TotalCount int      `json:"totalCount"`
	PodList    []string `json:"podList"`
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
	isInCluster, _ := getenv("IN_CLUSTER", "false")
	config.isInCluster, _ = strconv.ParseBool(isInCluster)
	return
}

func makeConfigFile(config Configuration) *rest.Config {
	if config.isInCluster {
		clientConfig, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		return clientConfig
	}

	clientConfig, err := clientcmd.BuildConfigFromFlags("", config.kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	return clientConfig
}

func buildKubeClient(clientConfig *rest.Config) *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func main() {
	config := configure()
	clientConfig := makeConfigFile(config)
	client := buildKubeClient(clientConfig)

	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	if config.format == "json" {
		names := make([]string, 0)
		for _, pod := range pods.Items {
			names = append(names, pod.GetName())
		}
		output := Output{len(pods.Items), names}
		result, err := json.Marshal(output)
		if err != nil {
			log.Fatalf("Error marshalling JSON, %f", err)
		}
		fmt.Println(string(result))
		return
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("Pod name=/%s\n", pod.GetName())
	}
}

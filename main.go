package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	client, err := kubernetes.NewForConfig(config.GetConfigOrDie())
	if err != nil {
		panic(err)
	}
	wc, err := client.CoreV1().Events("").Watch(context.Background(), metav1.ListOptions{Watch: true})
	if err != nil {
		panic(err)
	}
	resultChannel := wc.ResultChan()
	for {
		select {
		case event := <-resultChannel:
			eventBytes, _ := json.Marshal(event)
			fmt.Println(string(eventBytes))
		}
	}
}

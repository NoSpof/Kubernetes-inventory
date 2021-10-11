package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"nospof.cloud/common"
)

func GetNS(clientKate kubernetes.Clientset) []string {
	// Get Namespace

	ns, err := clientKate.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	var namespaces []string
	for _, namespace := range ns.Items {
		namespaces = append(namespaces, namespace.GetName())
	}
	return namespaces
}

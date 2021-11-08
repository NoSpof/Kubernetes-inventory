package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"nospof.cloud/common"
)

func GetNS(clientKate kubernetes.Clientset) []string {
	// Get Namespace
	excludeNs := []string{"kube-system"}
	ns, err := clientKate.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	var namespaces []string
	for _, namespace := range ns.Items {
		if !common.InArray(namespace.GetName(), excludeNs) {
			namespaces = append(namespaces, namespace.GetName())
		}
	}
	return namespaces
}

func GetResumeNs(clientKate kubernetes.Clientset, ns string) namespace {
	var nss namespace
	deploy, err := clientKate.AppsV1().Deployments(ns).List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	pods, err := clientKate.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	secrets, err := clientKate.CoreV1().Secrets(ns).List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	cms, err := clientKate.CoreV1().ConfigMaps(ns).List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	nss = namespace{
		Name:        ns,
		DeployCount: len(deploy.Items),
		PodCount:    len(pods.Items),
		SecretCount: len(secrets.Items),
		ConfigMap:   len(cms.Items),
	}
	return nss

}

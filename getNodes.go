package main

import (
	"context"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"nospof.cloud/common"
)

func getNodes(clientKate *kubernetes.Clientset) clusterCapacity {
	nodes, err := clientKate.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	var allCpu float64
	var allMemory float64

	for _, node := range nodes.Items {
		convertCpu, err := strconv.ParseFloat(node.Status.Allocatable.Cpu().AsDec().String(), 64)
		common.ExitIfError(err)
		convertMemory, err := strconv.ParseFloat(node.Status.Allocatable.Memory().AsDec().String(), 64)
		common.ExitIfError(err)
		allCpu = allCpu + convertCpu
		allMemory = allMemory + convertMemory
	}

	var responseObject = clusterCapacity{
		CPUCapacity:    allCpu,
		MemoryCapacity: allMemory,
	}
	return responseObject
}

package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"nospof.cloud/common"
)

func getNodes(clientKate *kubernetes.Clientset) []node {

	nodes, err := clientKate.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	var exportedNodes []node
	for _, nodeItem := range nodes.Items {
		nodetoExport := node{
			Name:             nodeItem.Name,
			UnSchedulable:    nodeItem.DeepCopy().DeepCopy().Spec.Unschedulable,
			Arch:             nodeItem.Status.NodeInfo.Architecture,
			ContainerRuntime: nodeItem.Status.NodeInfo.ContainerRuntimeVersion,
			KernelVersion:    nodeItem.Status.NodeInfo.KernelVersion,
			KubeProxyVersion: nodeItem.Status.NodeInfo.KubeProxyVersion,
			KubeletVersion:   nodeItem.Status.NodeInfo.KubeletVersion,
			OsImage:          nodeItem.Status.NodeInfo.OSImage,
			OsVersion:        nodeItem.Status.NodeInfo.OperatingSystem,
			
		}
		exportedNodes = append(exportedNodes, nodetoExport)
	}
	return exportedNodes
}

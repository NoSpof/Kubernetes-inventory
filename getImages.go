package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"nospof.cloud/common"
)

func getImages(clientKate *kubernetes.Clientset) []image {
	nodes, err := clientKate.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	var exportedImage []image
	for _, nodeItem := range nodes.Items {
		for _, imageItem := range nodeItem.Status.DeepCopy().Images {
			temporaryImage := image{
				Names:    imageItem.DeepCopy().Names[0],
				NodeName: nodeItem.GetName(),
				Size:     imageItem.SizeBytes / 1048576,
			}
			exportedImage = append(exportedImage, temporaryImage)
		}

	}
	return exportedImage
}

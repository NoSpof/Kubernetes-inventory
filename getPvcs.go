package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"nospof.cloud/common"
)

func getPvcs(clientKate *kubernetes.Clientset, namespace string) []pvc {
	pvcs, err := clientKate.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	common.ExitIfError(err)
	var pvcList []pvc
	for _, pvcItem := range pvcs.Items {
		pvcTemp := pvc{
			Name:      pvcItem.Name,
			Namespace: pvcItem.Namespace,
			Size:      pvcItem.Status.Capacity.Storage().Value() / 1048576,
		}
		pvcList = append(pvcList, pvcTemp)
	}
	return pvcList

}

package main

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"strconv"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"nospof.cloud/common"
)

func GetRqs(clientKate *kubernetes.Clientset, namespace string) resourcesquotas {
	var maxLimitCpu float64
	var maxLimitMemory float64
	var allocatedCpu float64
	var allocatedMemory float64

	rqs, err := clientKate.CoreV1().ResourceQuotas(namespace).List(context.TODO(), metav1.ListOptions{})
	log.Println("Parsing : " + namespace + " namespace in progress")
	common.ExitIfError(err)
	for _, rq := range rqs.Items {
		log.Println("Quotas found : " + rq.GetName())
		itemsCpu := rq.Spec.Hard.DeepCopy()["limits.cpu"]
		itemsMemory := rq.Spec.Hard.DeepCopy()["limits.memory"]

		// Convert to json and get the valid with trim
		jsonitems, err := json.MarshalIndent(itemsCpu, "", "    ")
		common.ExitIfError(err)
		convertCpuRqs, err := strconv.ParseFloat(strings.Trim(string(jsonitems), "\""), 64)
		common.ExitIfError(err)

		// Convert to json and get the valid with trim
		jsonitemsmemory, err := json.MarshalIndent(itemsMemory.AsApproximateFloat64(), "", "    ")
		common.ExitIfError(err)
		convertMemoryRqs, err := strconv.ParseFloat(strings.Trim(string(jsonitemsmemory), "\""), 64)
		common.ExitIfError(err)

		// Get the used memory and used Cpu
		usedMemory := rq.Status.Used.DeepCopy()["limits.memory"]
		usedCpu := rq.Status.Used.DeepCopy()["limits.cpu"]

		jsonUsedMemory, err := json.MarshalIndent(usedMemory.AsApproximateFloat64(), "", "    ")
		common.ExitIfError(err)
		jsonUsedCpu, err := json.MarshalIndent(usedCpu.AsApproximateFloat64(), "", "    ")
		common.ExitIfError(err)
		convertAllocatedCpu, err := strconv.ParseFloat(strings.Trim(string(jsonUsedCpu), "\""), 64)
		common.ExitIfError(err)

		convertAllocatedMemory, err := strconv.ParseFloat(strings.Trim(string(jsonUsedMemory), "\""), 64)
		common.ExitIfError(err)
		// Create the sum
		maxLimitCpu = maxLimitCpu + convertCpuRqs
		maxLimitMemory = maxLimitMemory + convertMemoryRqs
		allocatedCpu = allocatedCpu + convertAllocatedCpu
		allocatedMemory = allocatedMemory + convertAllocatedMemory

	}
	usageRateCpu := math.Round((allocatedCpu / maxLimitCpu) * 100)
	usageRateMemory := math.Round((allocatedMemory / maxLimitMemory) * 100)
	var responseObject = resourcesquotas{
		MaxLimitCpu:     common.FloatToString(maxLimitCpu),
		MaxLimitMemory:  common.FloatToString(maxLimitMemory),
		Namespace:       namespace,
		RateCpu:         common.FloatToString(usageRateCpu),
		RateMemory:      common.FloatToString(usageRateMemory),
		RequestedCpu:    common.FloatToString(allocatedCpu),
		RequestedMemory: common.FloatToString(allocatedMemory),
	}
	return responseObject
}

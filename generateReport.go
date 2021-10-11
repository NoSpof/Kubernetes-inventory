package main

import (
	"os"
	"strconv"

	"nospof.cloud/common"
)

func generateReport(nodes clusterCapacity, rqs []resourcesquotas, nodeList []node, imageList []image) {
	f, err := os.Create("./cluster.html")
	common.ExitIfError(err)
	defer f.Close()
	var content []byte
	header := []byte("<html><head><title>ClusterInfo Reporting</title></head><body>")
	footer := []byte("</body></html>")
	table := []byte("<table border='1'>")
	headerNode := []byte("<tr><th>Namespace</th><th>Cpu Request</th><th>Memory Request</th><th>Max Cpu</th><th>Max Memory</th><th>Rate Cpu</th><th>Rate Memory</th></tr>")
	headerNodePool := []byte("<tr><th>Name</th><th>Image Os</th><th>Kind</th><th>Container Runtime</th><th>Architecture</th><th>Kernel Version</th><th>Kubelet Version</th><th> kubeProxy Version</th><th>Unschedulable</th></tr>")
	headerImage := []byte("<tr><th>Image Name</th><th> Size (Mb) </th><th> Node name </th></tr>")
	endTable := []byte("</table >")
	titleCluster := []byte("<h2>Cluster Informations </h2>")
	titleNamespace := []byte("<h2>Ressource Informations</h2>")
	titleImage := []byte("<h2>Images Informations</h2>")
	var contentNode []byte
	var contentNodePool []byte
	var contentImage []byte
	for _, rq := range rqs {
		temporaryContent := []byte("<tr><td>" + rq.Namespace + "</td><td>" + rq.RequestedCpu + "</td><td>" + rq.RequestedMemory + "</td><td>" + rq.MaxLimitCpu + "</td><td>" + rq.MaxLimitMemory + "</td><td>" + rq.RateCpu + "</td><td>" + rq.RateMemory + "</td></tr>")
		contentNode = append(contentNode, temporaryContent...)
	}
	for _, node := range nodeList {
		temporaryNode := []byte("<tr><td>" + node.Name + "</td><td>" + node.OsImage + "</td><td>" + node.OsVersion + "</td><td>" + node.ContainerRuntime + "</td><td>" + node.Arch + "</td><td>" + node.KernelVersion + "</td><td>" + node.KubeletVersion + "</td><td>" + node.KubeProxyVersion + "</td><td>" + strconv.FormatBool(node.UnSchedulable) + "</td></tr>")
		contentNodePool = append(contentNodePool, temporaryNode...)
	}
	for _, image := range imageList {
		temporaryImage := []byte("<tr><td>" + image.Names + "</td><td>" + strconv.FormatInt(image.Size, 10) + "</td><td>" + image.NodeName + "</td></tr>")
		contentImage = append(contentImage, temporaryImage...)
	}
	// Write on outût html
	content = append(content, header...)
	content = append(content, titleCluster...)
	content = append(content, table...)
	content = append(content, headerNodePool...)
	content = append(content, contentNodePool...)
	content = append(content, endTable...)
	content = append(content, titleNamespace...)
	content = append(content, table...)
	content = append(content, headerNode...)
	content = append(content, contentNode...)
	content = append(content, endTable...)

	content = append(content, titleImage...)
	content = append(content, table...)
	content = append(content, headerImage...)
	content = append(content, contentImage...)
	content = append(content, endTable...)
	content = append(content, footer...)
	_, err = f.Write(content)
	common.ExitIfError(err)
}

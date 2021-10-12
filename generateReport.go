package main

import (
	"os"
	"strconv"

	"nospof.cloud/common"
)

func generateReport(nodes clusterCapacity, rqs []resourcesquotas, nodeList []node, imageList []image, pvcList []pvc, nsList []namespace) {
	f, err := os.Create("./cluster.html")
	common.ExitIfError(err)
	defer f.Close()
	var content []byte
	header := []byte("<html><head><title>ClusterInfo Reporting</title><link href='https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css' rel='stylesheet' integrity='sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC' crossorigin='anonymous'></head><body>")
	js := []byte("<script src='https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js' integrity='sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM' crossorigin='anonymous'></script>")
	footer := []byte("</body></html>")
	table := []byte("<table class='table'>")
	headerNamespace := []byte("<thead><tr><th scope='col'>Name</th><th scope='col'>Count Deploy</th><th scope='col'>Count Pods</th><th scope='col'>Count Secrets</th><th scope='col'>Count Configmap</th></tr></thead><tbody>")
	headerNode := []byte("<thead><tr><th scope='col'>Namespace</th><th scope='col'>Cpu Request</th><th scope='col'>Memory Request</th><th scope='col'>Max Cpu</th><th scope='col'>Max Memory</th><th scope='col'>Rate Cpu</th><th scope='col'>Rate Memory</th></tr></thead><tbody>")
	headerNodePool := []byte("<thead><tr><th scope='col'>Name</th><th scope='col'>Image Os</th><th scope='col'>Kind</th><th scope='col'>Container Runtime</th><th scope='col'>Architecture</th><th scope='col'>Kernel Version</th><th scope='col'>Kubelet Version</th><th scope='col'> kubeProxy Version</th><th scope='col'>Unschedulable</th></tr></thead><tbody>")
	headerImage := []byte("<thead><tr><th scope='col'>Image Name</th><th scope='col'> Size (Mb) </th><th scope='col'> Node name </th></tr></thead><tbody>")
	headerPvc := []byte("<thead><tr><th scope='col'>Pvc Name</th><th scope='col'>Pvc Namespace</th><th scope='col'>Pvc Size</th></tr></thead><tbody>")
	endTable := []byte("</tbody></table >")
	titleNamespace := []byte("<h2>Namespace Informations : </h2>")
	titleCluster := []byte("<h2>Cluster Informations :  </h2>")
	titleRessource := []byte("<h2>Ressource Informations : </h2>")
	titleImage := []byte("<h2>Images Informations : </h2>")
	titlePvc := []byte("<h2>Pvc Informations : </h2>")
	var contentNode []byte
	var contentNodePool []byte
	var contentImage []byte
	var contentPvc []byte
	var contentNamespace []byte
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
	for _, pvc := range pvcList {
		temporaryPvc := []byte("<tr><td>" + pvc.Name + "</td><td>" + pvc.Namespace + "</td><td>" + strconv.FormatInt(pvc.Size, 10) + "</td></tr>")
		contentPvc = append(contentPvc, temporaryPvc...)
	}
	for _, ns := range nsList {
		temporaryNs := []byte("<tr><td>" + ns.Name + "</td><td>" + strconv.Itoa(ns.DeployCount) + "</td><td>" + strconv.Itoa(ns.PodCount) + "</td><td>" + strconv.Itoa(ns.SecretCount) + "</td><td>" + strconv.Itoa(ns.ConfigMap) + "</td></tr>")
		contentNamespace = append(contentNamespace, temporaryNs...)
	}

	// Write on output html
	content = append(content, header...)
	content = append(content, js...)
	content = append(content, titleCluster...)
	content = append(content, table...)
	content = append(content, headerNodePool...)
	content = append(content, contentNodePool...)
	content = append(content, endTable...)
	content = append(content, titleNamespace...)
	content = append(content, table...)
	content = append(content, headerNamespace...)
	content = append(content, contentNamespace...)
	content = append(content, endTable...)
	content = append(content, titleRessource...)
	content = append(content, table...)
	content = append(content, headerNode...)
	content = append(content, contentNode...)
	content = append(content, endTable...)
	content = append(content, titleImage...)
	content = append(content, table...)
	content = append(content, headerImage...)
	content = append(content, contentImage...)
	content = append(content, endTable...)
	content = append(content, titlePvc...)
	content = append(content, table...)
	content = append(content, headerPvc...)
	content = append(content, contentPvc...)
	content = append(content, endTable...)
	content = append(content, footer...)
	_, err = f.Write(content)
	common.ExitIfError(err)
}

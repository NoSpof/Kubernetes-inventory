package main

func main() {
	clientKate := connectKubernetes()
	clusterCapacity := getClusterCapacity(clientKate)
	nsArray := GetNS(*clientKate)
	var rqsArray []resourcesquotas
	var pvcList []pvc
	var nsList []namespace
	for _, ns := range nsArray {
		responserq := GetRqs(clientKate, ns)
		rqsArray = append(rqsArray, responserq)
		responsePvcs := getPvcs(clientKate, ns)
		pvcList = append(pvcList, responsePvcs...)
		responseNss := GetResumeNs(*clientKate, ns)
		nsList = append(nsList, responseNss)

	}
	nodesList := getNodes(clientKate)
	//imageList := getImages(clientKate)
	var imageList []image
	generateReport(clusterCapacity, rqsArray, nodesList, imageList, pvcList, nsList)
}

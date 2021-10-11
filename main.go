package main

func main() {
	clientKate := connectKubernetes()
	clusterCapacity := getClusterCapacity(clientKate)
	nsArray := GetNS(*clientKate)
	var rqsArray []resourcesquotas
	for _, ns := range nsArray {
		responserq := GetRqs(clientKate, ns)
		rqsArray = append(rqsArray, responserq)
	}
	nodesList := getNodes(clientKate)
	imageList := getImages(clientKate)
	generateReport(clusterCapacity, rqsArray, nodesList, imageList)
}

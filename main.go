package main

func main() {
	clientKate := connectKubernetes()
	nodes := getNodes(clientKate)
	nsArray := GetNS(*clientKate)
	var rqsArray []resourcesquotas
	for _, ns := range nsArray {
		responserq := GetRqs(clientKate, ns)
		rqsArray = append(rqsArray, responserq)
	}
	generateReport(nodes, rqsArray)
}

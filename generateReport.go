package main

import (
	"os"

	"nospof.cloud/common"
)

func generateReport(nodes clusterCapacity, rqs []resourcesquotas) {
	f, err := os.Create("./cluster2.html")
	common.ExitIfError(err)
	defer f.Close()
	header := "<html><head><title>ClusterInfo Reporting</title></head><body>"
	footer := "</body></html>"
	table := "<table border='1'>"
	headerNode := []byte("<tr><th>Namespace</th><th>Cpu Request</th><th>Memory Request</th><th>Max Cpu</th><th>Max Memory</th><th>Rate Cpu</th><th>Rate Memory</th></tr>")
	endTable := "</table >"
	_, err = f.WriteString(header)
	common.ExitIfError(err)
	var contentNode []byte
	_, err = f.WriteString("<h2>Cluster Informations</h2>")
	common.ExitIfError(err)
	_, err = f.WriteString("<h2>Ressource Informations</h2>")
	common.ExitIfError(err)
	_, err = f.WriteString(table)
	common.ExitIfError(err)
	for _, rq := range rqs {
		temporaryContent := []byte("<tr><td>" + rq.Namespace + "</td><td>" + rq.RequestedCpu + "</td><td>" + rq.RequestedMemory + "</td><td>" + rq.MaxLimitCpu + "</td><td>" + rq.MaxLimitMemory + "</td><td>" + rq.RateCpu + "</td><td>" + rq.RateMemory + "</td></tr>")
		contentNode = append(contentNode, temporaryContent...)
	}
	_, err = f.Write(headerNode)
	common.ExitIfError(err)
	_, err = f.Write(contentNode)
	common.ExitIfError(err)
	_, err = f.WriteString(endTable)
	common.ExitIfError(err)
	_, err = f.WriteString(footer)
	common.ExitIfError(err)
}

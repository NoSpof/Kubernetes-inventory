package main

type clusterCapacity struct {
	CPUCapacity    float64 `json:"cpuCapacity"`
	MemoryCapacity float64 `json:"memoryCapacity"`
}

type resourcesquotas struct {
	MaxLimitCpu     string
	MaxLimitMemory  string
	RequestedCpu    string
	RequestedMemory string
	Namespace       string
	RateCpu         string
	RateMemory      string
}
type rqsArray []resourcesquotas

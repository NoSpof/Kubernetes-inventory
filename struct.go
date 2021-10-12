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
type node struct {
	Name             string
	UnSchedulable    bool
	KernelVersion    string
	OsImage          string
	ContainerRuntime string
	KubeletVersion   string
	KubeProxyVersion string
	OsVersion        string
	Arch             string
}

type image struct {
	Names    string
	NodeName string
	Size     int64
}

type pvc struct {
	Name      string
	Size      int64
	Namespace string
}
type namespace struct {
	Name        string
	DeployCount int
	PodCount    int
	SecretCount int
	ConfigMap   int
}

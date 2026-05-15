package model


type WorkloadMetric struct {
	Deployment     string `json:"deployment"`
	CPURequest     int    `json:"cpu_request"`
	CPUUsageAvg    int    `json:"cpu_usage_avg"`
	MemoryRequest  int    `json:"memory_request"`
	MemoryUsageAvg int    `json:"memory_usage_avg"`
}

type Recommendation struct {
	Deployment        string `json:"deployment"`
	RecommendedCPU    int    `json:"recommended_cpu"`
	RecommendedMemory int    `json:"recommended_memory"`
	Reason            string `json:"reason"`
}
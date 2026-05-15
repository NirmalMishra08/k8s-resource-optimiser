package optimizer

import (
	"k8s-resource-optimizer/internal/model"
	"k8s-resource-optimizer/internal/utils"
)

const (
	CPUBuffer    = 1.5
	MemoryBuffer = 1.3
)

func GenerateRecommendation(
	workload model.WorkloadMetric,
) *model.Recommendation {

	cpuTarget := utils.RoundCPU(
		int(float64(workload.CPUUsageAvg) * CPUBuffer),
	)

	memoryTarget := utils.RoundMemory(
		int(float64(workload.MemoryUsageAvg) * MemoryBuffer),
	)

	cpuOverprovisioned :=
		workload.CPURequest > int(float64(cpuTarget)*1.2)

	memoryOverprovisioned :=
		workload.MemoryRequest > int(float64(memoryTarget)*1.2)

	if cpuOverprovisioned || memoryOverprovisioned {
		return &model.Recommendation{
			Deployment:        workload.Deployment,
			RecommendedCPU:    cpuTarget,
			RecommendedMemory: memoryTarget,
			Reason:            "Average usage significantly below requested resources",
		}
	}

	return nil
}

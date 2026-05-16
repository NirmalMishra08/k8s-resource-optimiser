package tests

import (
	"testing"

	"k8s-resource-optimizer/internal/model"
	"k8s-resource-optimizer/internal/optimizer"
)

func TestGenerateRecommendation_Overprovisioned(
	t *testing.T,
) {

	workload := model.WorkloadMetric{
		Deployment:     "api-service",
		CPURequest:     1000,
		CPUUsageAvg:    180,
		MemoryRequest:  2048,
		MemoryUsageAvg: 700,
	}

	rec := optimizer.GenerateRecommendation(
		workload,
	)

	if rec == nil {
		t.Fatal("expected recommendation")
	}

	if rec.RecommendedCPU != 300 {
		t.Errorf(
			"expected CPU 300 got %d",
			rec.RecommendedCPU,
		)
	}

	if rec.RecommendedMemory != 1024 {
		t.Errorf(
			"expected memory 1024 got %d",
			rec.RecommendedMemory,
		)
	}
}

func TestGenerateRecommendation_HealthyWorkload(
	t *testing.T,
) {

	workload := model.WorkloadMetric{
		Deployment:     "worker-service",
		CPURequest:     500,
		CPUUsageAvg:    450,
		MemoryRequest:  1024,
		MemoryUsageAvg: 900,
	}

	rec := optimizer.GenerateRecommendation(
		workload,
	)

	if rec != nil {
		t.Fatal(
			"expected no recommendation",
		)
	}
}

func TestGenerateRecommendation_MinCPUFloor(
	t *testing.T,
) {

	workload := model.WorkloadMetric{
		Deployment:     "tiny-service",
		CPURequest:     1000,
		CPUUsageAvg:    10,
		MemoryRequest:  2048,
		MemoryUsageAvg: 100,
	}

	rec := optimizer.GenerateRecommendation(
		workload,
	)

	if rec.RecommendedCPU < 100 {
		t.Errorf(
			"cpu recommendation below minimum",
		)
	}
}
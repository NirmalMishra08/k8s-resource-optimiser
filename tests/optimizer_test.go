package tests

import (
	"testing"

	"k8s-resource-optimizer/internal/model"
	"k8s-resource-optimizer/internal/optimizer"
)

func TestGenerateRecommendation(
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
}

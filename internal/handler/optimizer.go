package handler

import (
	"encoding/json"
	"k8s-resource-optimizer/internal/metrics"
	"k8s-resource-optimizer/internal/model"
	"k8s-resource-optimizer/internal/optimizer"
	"net/http"
)

func OptimizerHandler(w http.ResponseWriter, r *http.Request){
	var workload []model.WorkloadMetric

	err:= json.NewDecoder(r.Body).Decode(&workload)
	if err != nil {
		http.Error(w, "invalid payload" , http.StatusBadRequest)
		return
	}

	metrics.RecommendationsGenerated.Inc()

	var recommedations []model.Recommendation

	for _, workload := range workload {
		 rec:= optimizer.GenerateRecommendation(workload)
		 if rec != nil{
			recommedations = append(recommedations, *rec)
		 }
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(recommedations)
}
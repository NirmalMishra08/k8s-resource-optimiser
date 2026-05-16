package main

import (
	"k8s-resource-optimizer/internal/handler"
	"k8s-resource-optimizer/internal/metrics"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := chi.NewRouter()

	r.Post("/optimise", handler.OptimizerHandler)
	log.Println("Server running on part :8080")

	prometheus.MustRegister(
		metrics.RecommendationsGenerated,
	)
	r.Handle(
		"/metrics",
		promhttp.Handler(),
	)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

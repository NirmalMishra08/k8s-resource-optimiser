package metrics

import "github.com/prometheus/client_golang/prometheus"

var RecommendationsGenerated = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "recommendations_generated_total",
		Help: "Total recommendations generated",
	},
)

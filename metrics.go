package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define the metrics
var (
	crashLoopCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "docker_container_crashloop_count",
			Help: "Number of times a container has entered a crash loop.",
		},
		[]string{"container_id", "container_name"}, // Labels to differentiate containers
	)
)

func initMetrics() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(crashLoopCounter)
}

// UpdateCrashLoopMetric updates the crash loop count metric for a given container
func UpdateCrashLoopMetric(containerID, containerName string) {
	// Increment the counter for the given container
	crashLoopCounter.WithLabelValues(containerID, containerName).Inc()
}

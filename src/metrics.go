package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define the metrics
var (
	// Existing crash loop counter
	crashLoopCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "docker_container_crashloop_count",
			Help: "Number of times a container has entered a crash loop, labeled by exit code.",
		},
		[]string{"container_id", "container_name", "exit_code"},
	)

	// New counter for all container stops/restarts
	containerStopRestartCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "docker_container_stop_restart_count",
			Help: "Count of all container stops/restarts, labeled by exit code.",
		},
		[]string{"container_id", "container_name", "exit_code"},
	)
)

func initMetrics() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(crashLoopCounter)
	prometheus.MustRegister(containerStopRestartCounter)
}

// UpdateCrashLoopMetric updates the crash loop count metric for a given container with the exit code label
func UpdateCrashLoopMetric(containerID, containerName string, exitCode string) {
	crashLoopCounter.WithLabelValues(containerID, containerName, exitCode).Inc()
}

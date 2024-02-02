package main

import (
	"log"
	"net/http"

	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error initializing Docker client: %v", err)
	}

	// Define Prometheus metrics
	initMetrics()

	// Monitor Docker events in the background
	go monitorDockerEvents(cli)

	// Setup HTTP server for Prometheus scraping
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

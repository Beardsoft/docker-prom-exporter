package main

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// Struct for holding container state information
type ContainerState struct {
	restartCount   int
	lastState      string
	lastUpdateTime time.Time // Make sure this line is added
}

// Map for tracking container states by their IDs
var containerStates = make(map[string]*ContainerState)

func monitorDockerEvents(cli *client.Client) {
	// Define a filter to listen for start and die events
	eventFilter := filters.NewArgs()
	eventFilter.Add("type", "container")
	eventFilter.Add("event", "start")
	eventFilter.Add("event", "die")

	options := types.EventsOptions{
		Filters: eventFilter,
		Since:   time.Now().Format(time.RFC3339),
	}

	// Listen for events
	events, errors := cli.Events(context.Background(), options)
	for {
		select {
		case event := <-events:
			handleDockerEvent(event)
		case err := <-errors:
			if err != nil {
				log.Fatalf("Error while listening to Docker events: %v", err)
			}
		}
	}
}

func handleDockerEvent(event events.Message) {
	containerID := event.Actor.ID
	containerName := event.Actor.Attributes["name"]
	exitCode := event.Actor.Attributes["exitCode"] // Extracting the exit code

	// Update the metric for all container stops/restarts
	containerStopRestartCounter.WithLabelValues(containerID, containerName, exitCode).Inc()

	if event.Action == "die" {
		if state, exists := containerStates[containerID]; exists {
			if state.lastState == "started" && time.Since(state.lastUpdateTime).Minutes() < 5 {
				// This is a simplistic way to identify a crash-loop
				log.Printf("Container %s (%s) is potentially crash-looping with exit code: %s", containerName, containerID, exitCode)
				UpdateCrashLoopMetric(containerID, containerName, exitCode)
			}
			state.restartCount++
			state.lastState = "died"
		} else {
			containerStates[containerID] = &ContainerState{
				restartCount:   1,
				lastState:      "died",
				lastUpdateTime: time.Now(),
			}
		}
	} else if event.Action == "start" {
		if state, exists := containerStates[event.Actor.ID]; exists {
			state.lastState = "started"
			state.lastUpdateTime = time.Now()
		}
	}
}

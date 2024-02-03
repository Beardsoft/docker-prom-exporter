# Docker Container Monitoring Application

## Overview

This application monitors Docker containers, tracking crash-looping behaviors and general container stop/restart events. It exposes metrics compatible with Prometheus, allowing for real-time monitoring and alerting based on container status and exit codes.

## Features

- **Crash Loop Detection:** Identifies containers that are repeatedly crashing and restarting within a short timeframe.
- **Exit Code Monitoring:** Tracks the exit codes of stopped or restarted containers, providing insights into why containers may have stopped.
- **Prometheus Integration:** Exposes metrics in a Prometheus-compatible format for easy scraping, visualization, and alerting.

## Prerequisites

- Docker: Ensure Docker is installed and running on your system.
- Go: This application is written in Go. Make sure Go is installed on your system.
- Prometheus: For scraping metrics exposed by this application.
- Grafana (optional): For visualizing the metrics.

## Setup and Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/maestroi/docker-prom-exporter.git
   cd docker-monitoring-app
    ```

2. **Build:**
```bash
go build -o docker-monitor
```

3. **Run:**
```bash
./docker-monitor
```

## Setup with Docker
docker build -t docker-monitor .
docker run -p 8080:8088 -v /var/run/docker.sock:/var/run/docker.sock yourusername/docker-monitor

## Confguration
No additional configuration is needed to start monitoring your Docker containers. However, ensure Prometheus is configured to scrape metrics from the application's exposed endpoint (default is :8080/metrics).

## Prometheus Scrape Configuration
Add the following job to your Prometheus configuration file:

```yaml
scrape_configs:
  - job_name: 'docker-monitor'
    static_configs:
      - targets: ['localhost:8080']
```
Adjust the target if your application is running on a different host or port.

## Metrics Exposed
docker_container_crashloop_count: Counts the number of times a container has entered a crash loop, labeled by the container ID, name, and exit code.
docker_container_stop_restart_count: Counts all container stops/restarts, labeled by the container ID, name, and exit code.

## Viewing Metrics
Once Prometheus is scraping metrics from the application, you can view these metrics directly in Prometheus or use Grafana to create dashboards for visualization and alerting.

## Contributing
Contributions to this project are welcome! Please submit issues and pull requests with any enhancements, bug fixes, or suggestions.

## License
Specify your license here or indicate if the project is open-source and freely available for use.

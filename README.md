# InfraWatch

InfraWatch is a DevOps simulation project that demonstrates how to deploy and maintain a lightweight Go API service using Docker, Nginx reverse proxy, Jenkins pipeline, health check automation, and backup scripting.

## CI/CD

This project includes two CI/CD approaches:

- Jenkins Pipeline for simulated deployment automation
- GitHub Actions for automated testing and Docker image build validation

## Tech Stack

- Go
- Docker
- Docker Compose
- Nginx
- Jenkins
- Shell Script
- Linux-style operational workflow

## Features

- Go REST API service
- `/health` endpoint for service monitoring
- `/version` endpoint for release tracking
- Dockerized application
- Nginx reverse proxy
- Jenkins CI/CD pipeline
- Automated health check script
- Automated backup script
- Deployment documentation
- GitHub Actions CI workflow
- Go unit tests

## Service Endpoints

| Endpoint   | Description                     |
| ---------- | ------------------------------- |
| `/`        | Basic service status            |
| `/health`  | Health check response           |
| `/version` | Application version info        |
| `/metrics` | Basic simulated service metrics |

## Environment Variables

Create a local environment file from the example:

```bash
cp .env.example .env
```

Example configuration:

```env
APP_NAME=InfraWatch
APP_ENV=development
APP_PORT=8080
NGINX_PORT=8082
```

## Run Locally

```bash
docker compose up -d --build
```

Check service through Nginx:

```bash
curl http://localhost:8082/health
```

## Run Health Check

```bash
./scripts/health-check.sh
```

## Run Backup

```bash
./scripts/backup.sh
```

## Check Metrics

```bash
curl http://localhost:8082/metrics
```

## Project Goal

This project was built to simulate real DevOps responsibilities such as application deployment, reverse proxy configuration, CI/CD automation, monitoring, backup, and operational documentation.

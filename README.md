# InfraWatch

InfraWatch is a DevOps simulation project that demonstrates how to deploy and maintain a lightweight Go API service using Docker, Nginx reverse proxy, Jenkins pipeline, health check automation, and backup scripting.

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

## Service Endpoints

| Endpoint   | Description              |
| ---------- | ------------------------ |
| `/`        | Basic service status     |
| `/health`  | Health check response    |
| `/version` | Application version info |

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

## Project Goal

This project was built to simulate real DevOps responsibilities such as application deployment, reverse proxy configuration, CI/CD automation, monitoring, backup, and operational documentation.

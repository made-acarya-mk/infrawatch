# InfraWatch Architecture

## Overview

InfraWatch is a lightweight service deployment project designed to simulate a real DevOps workflow.

The system consists of:

- Go API service
- Nginx reverse proxy
- Docker Compose orchestration
- Jenkins deployment pipeline
- GitHub Actions CI workflow
- Shell scripts for monitoring, backup, logs, and status checking

## Architecture Diagram

```txt
User / Client
     |
     v
Nginx Reverse Proxy
     |
     v
InfraWatch Go API
     |
     v
Health, Version, and Metrics Endpoints
```

## Request Flow

1. User sends request to Nginx through port 8082.
2. Nginx forwards the request to the Go API container.
3. Go API processes the request.
4. Response is returned back through Nginx.

## Container Services

| Service          | Purpose                                      |
| ---------------- | -------------------------------------------- |
| infrawatch-api   | Runs the Go API service                      |
| infrawatch-nginx | Acts as reverse proxy and public entry point |

## Operational Flow

```txt
Code Change
↓
CI Test
↓
Docker Build
↓
Deployment
↓
Health Check
↓
Backup
↓
Monitoring / Incident Response
```

# Deployment SOP

## Objective

This document explains the standard deployment process for the InfraWatch service.

## Prerequisites

- Docker installed
- Docker Compose installed
- Git installed
- Port 8082 available on host machine

## Deployment Steps

### 1. Pull latest code

```bash
git pull origin main
```

### 2. Build and start services

```bash
docker compose up -d --build
```

### 3. Verify running containers

```bash
docker ps
```

### 4. Run health check

```bash
./scripts/health-check.sh
```

### 5. Validate service manually

```bash
curl http://localhost:8082/health
```

## Rollback Plan

If deployment fails:

```bash
docker compose down
```

Then restore from the latest backup file in the backups/ directory.

## Success Criteria

Deployment is successful if:

- API container is running
- Nginx container is running
- /health returns healthy status
- Health check log records OK status

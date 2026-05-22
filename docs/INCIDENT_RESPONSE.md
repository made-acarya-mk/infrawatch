# Incident Response Runbook

## Objective

This document provides a basic incident response flow when the InfraWatch service is unreachable or unhealthy.

## Incident Scenario

Service health check fails or `/health` endpoint does not return a healthy response.

## Severity Level

| Level  | Description                        |
| ------ | ---------------------------------- |
| Low    | Temporary slow response            |
| Medium | Service intermittently unavailable |
| High   | Service completely unreachable     |

## Troubleshooting Steps

### 1. Check running containers

```bash
docker ps
```

### 2. Check container logs

```bash
docker logs infrawatch-api
docker logs infrawatch-nginx
```

### 3. Check Nginx route

```bash
curl http://localhost:8082/health
```

### 4. Check API directly inside Docker network

```bash
docker exec -it infrawatch-nginx wget -qO- http://infrawatch-api:8080/health
```

### 5. Restart services

```bash
docker compose restart
```

### 6. Rebuild services if needed

```bash
docker compose down
docker compose up -d --build
```

## Recovery Criteria

Incident is resolved when:

- API container is running
- Nginx container is running
- `/health` returns healthy
- Health check script logs OK status

## Escalation Notes

If the issue cannot be resolved at this level, collect:

- Container logs
- Health check logs
- Recent deployment changes
- Error screenshots

Then escalate to senior engineer or infrastructure team.

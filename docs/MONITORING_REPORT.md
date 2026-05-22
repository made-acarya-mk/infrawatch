# Monitoring Report

## Service Information

| Item              | Value                        |
| ----------------- | ---------------------------- |
| Service Name      | InfraWatch API               |
| Environment       | Local / Simulated Production |
| Reverse Proxy     | Nginx                        |
| Container Runtime | Docker                       |

## Health Check Target

```bash
http://localhost:8082/health
```

## Monitoring Method

The service is monitored using a shell script:

```bash
./scripts/health-check.sh
```

The script sends an HTTP request to the /health endpoint and writes the result into:

```bash
logs/health.log
```

## Sample Log

```txt
2026-05-22 10:00:00 | OK | InfraWatch API is healthy
2026-05-22 10:05:00 | FAILED | InfraWatch API is not reachable
```

## Alert Condition

An incident should be investigated when:

- Health check returns FAILED
- API container is stopped
- Nginx container is stopped
- `/health` endpoint does not respond

## Basic Action

When failure is detected:

```bash
docker ps
docker logs infrawatch-api
docker logs infrawatch-nginx
docker compose restart
```

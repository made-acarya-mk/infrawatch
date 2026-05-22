# API Reference

## Base URL

```bash
http://localhost:8082
```

## Endpoints

### GET /

Returns basic service status.

```bash
curl http://localhost:8082/
```

Example response:

```txt
InfraWatch API is running
```

### GET /health

Returns service health status.

```bash
curl http://localhost:8082/health
```

Example response:

```txt
{
  "service": "infrawatch-api",
  "status": "healthy",
  "version": "1.0.0",
  "timestamp": "2026-05-22T10:00:00Z"
}
```

### GET /version

Returns application version.

```bash
curl http://localhost:8082/version
```

Example response:

```txt
{
  "version": "1.0.0",
  "environment": "development"
}
```

### GET /metrics

Returns basic simulated service metrics.

```bash
curl http://localhost:8082/metrics
```

Example response:

```txt
{
  "cpu_usage": "simulated",
  "memory_usage": "simulated",
  "request_status": "ok",
  "service": "infrawatch-api",
  "timestamp": "2026-05-22T10:00:00Z",
  "uptime_status": "running"
}
```

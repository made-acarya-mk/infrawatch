# Contributing Guide

## Branching Strategy

Use the following branch naming convention:

```txt
feature/short-description
fix/short-description
docs/short-description
```

## Development Flow

1. Create a new branch.
2. Make changes locally.
3. Run tests.
4. Build Docker services.
5. Submit a pull request.

## Local Validation

Before submitting changes, run:

```bash
cd app
go test ./...
cd ..
docker compose up -d --build
curl http://localhost:8082/health
```

## Commit Message Examples

```txt
Add health check endpoint
Fix Nginx reverse proxy config
Update deployment SOP
```

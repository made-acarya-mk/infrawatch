# Backup and Recovery SOP

## Objective

This document explains how to create and restore backups for the InfraWatch service.

## Backup Scope

The backup script includes:

- Application source code
- Nginx configuration
- Shell scripts
- Docker Compose file
- README documentation

## Create Backup

Run:

```bash
./scripts/backup.sh
```

## Verify Backup

Check generated backup files:

```bash
ls -lh backups
```

## Restore Backup

Extract the selected backup file:

```bash
tar -xzf backups/infrawatch-backup-YYYY-MM-DD_HH-MM-SS.tar.gz
```

## Restart Service After Restore

```bash
docker compose down
docker compose up -d --build
```

## Validate Recovery

```bash
curl http://localhost:8082/health
./scripts/health-check.sh
```

## Recovery Success Criteria

Recovery is successful if:

- Backup file can be extracted
- Docker services can start normally
- `/health` endpoint returns healthy
- Health check log records OK status

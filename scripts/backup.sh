#!/bin/bash

DATE=$(date '+%Y-%m-%d_%H-%M-%S')
BACKUP_DIR="./backups"
BACKUP_FILE="$BACKUP_DIR/infrawatch-backup-$DATE.tar.gz"

mkdir -p "$BACKUP_DIR"

tar -czf "$BACKUP_FILE" \
  app \
  nginx \
  scripts \
  docker-compose.yml \
  README.md

echo "Backup created: $BACKUP_FILE"
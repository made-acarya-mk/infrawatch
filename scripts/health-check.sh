#!/bin/bash

URL="http://localhost:8082/health"
LOG_FILE="./logs/health.log"

mkdir -p ./logs

if curl -fsS "$URL" > /dev/null; then
  echo "$(date '+%Y-%m-%d %H:%M:%S') | OK | InfraWatch API is healthy" >> "$LOG_FILE"
else
  echo "$(date '+%Y-%m-%d %H:%M:%S') | FAILED | InfraWatch API is not reachable" >> "$LOG_FILE"
fi
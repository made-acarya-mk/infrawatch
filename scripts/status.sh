#!/bin/bash

echo "InfraWatch Service Status"
echo "========================="
echo ""

echo "Docker containers:"
docker ps --filter "name=infrawatch" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

echo ""
echo "Health endpoint:"
curl -s http://localhost:8082/health || echo "Health endpoint unreachable"

echo ""
echo ""
echo "Metrics endpoint:"
curl -s http://localhost:8082/metrics || echo "Metrics endpoint unreachable"
echo ""
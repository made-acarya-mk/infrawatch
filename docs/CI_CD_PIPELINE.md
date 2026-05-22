# CI/CD Pipeline Documentation

## Objective

This document explains the Jenkins pipeline used to build, deploy, validate, and back up the InfraWatch service.

## Pipeline Tool

Jenkins is used as the CI/CD automation tool.

## Pipeline Stages

| Stage                             | Description                                  |
| --------------------------------- | -------------------------------------------- |
| Checkout                          | Pulls the latest source code from repository |
| Run Go Tests                      | Runs unit tests for Go handlers              |
| Build Docker Image                | Builds the Go API Docker image               |
| Run Container with Docker Compose | Starts the API and Nginx services            |
| Health Check                      | Runs automated service health validation     |
| Backup                            | Creates backup after deployment              |

## Jenkinsfile Location

```bash
Jenkinsfile
```

## Pipeline Flow

```txt
Code Push
   ↓
Jenkins Pipeline Trigger
   ↓
Build Docker Image
   ↓
Deploy with Docker Compose
   ↓
Run Health Check
   ↓
Create Backup
```

## Success Criteria

Pipeline is successful if:

- Go tests pass successfully
- Docker image builds successfully
- Containers run successfully
- `/health` endpoint returns healthy
- Backup file is created

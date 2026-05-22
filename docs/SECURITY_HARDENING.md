# Security Hardening Notes

## Objective

This document explains basic security practices applied to the InfraWatch deployment.

## Nginx Security Headers

The Nginx reverse proxy adds basic security headers:

- `X-Frame-Options`
- `X-Content-Type-Options`
- `X-XSS-Protection`

## Container Security Practices

- Application runs inside Docker container
- Only Nginx is exposed to host machine
- API service is only accessible inside Docker network
- Containers use restart policy `unless-stopped`

## Port Exposure

| Service        | Internal Port | Exposed Port |
| -------------- | ------------: | -----------: |
| InfraWatch API |          8080 |  Not exposed |
| Nginx          |            80 |         8082 |

## Recommended Future Improvements

- Add HTTPS with SSL certificate
- Add rate limiting in Nginx
- Add non-root Docker user
- Add vulnerability scanning
- Store secrets in environment variables

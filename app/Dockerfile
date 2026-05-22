FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o infrawatch-api main.go

FROM alpine:latest

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/infrawatch-api .

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./infrawatch-api"]
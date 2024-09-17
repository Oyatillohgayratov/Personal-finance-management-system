FROM golang:1.23.0-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /budget-service ./cmd/main.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /budget-service /budget-service
EXPOSE 9001
CMD ["/budget-service"]

FROM golang:latest AS builder

WORKDIR /opt/
ADD go.mod ./
COPY ./src ./src
RUN go mod download && \
    GOARCH=amd64 GOOS=linux go build -o /opt/main /opt/src/main.go

FROM alpine:latest
COPY --from=builder /opt/main /opt/main
RUN chmod +x /opt/main

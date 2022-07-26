FROM golang:latest AS builder

WORKDIR /opt/
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ADD go.mod ./
COPY ./src ./src
RUN go mod download && \
    go build -o /opt/main /opt/src/main.go

FROM alpine:latest
COPY --from=builder /opt/main /opt/main
RUN chmod +x /opt/main

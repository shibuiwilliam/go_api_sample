FROM golang:latest AS builder

WORKDIR /opt/
ADD go.mod ./
ADD go.sum ./
COPY ./src ./src
RUN go mod download && \
    go build -o /opt/main /opt/src/main.go

FROM golang:latest
COPY --from=builder /opt/main /opt/main
RUN chmod +x /opt/main

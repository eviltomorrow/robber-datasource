# Build the manager binary
FROM golang:1.17.4-buster AS builder

LABEL maintainer="eviltomorrow@163.com"

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct
ENV WORKSPACE=/workspace/
ENV CGO_ENABLED=1

WORKDIR $WORKSPACE
ADD . .

# Build
RUN make build

# Run
FROM alpine:3.12
# Copy binary file
COPY --from=builder /tmp/robber-datasource /usr/bin/
COPY --from=builder /workspace/liarsa.com/robber-datasource/config/config.toml /etc/robber-datasource/config.toml

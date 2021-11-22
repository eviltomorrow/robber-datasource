# Build the manager binary
FROM golang:1.15.8-buster AS builder

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct
ENV WORKSPACE=/workspace/liarsa.com/robber-datasource

WORKDIR $WORKSPACE
ADD . .

# Build
RUN mkdir -p bin && go build -ldflags "-X main.GitSha=`git rev-parse --short HEAD` -X main.GitTag=`git describe --tags --always` -X main.GitBranch=`git rev-parse --abbrev-ref HEAD` -X main.BuildTime=`date +%FT%T%z` -s -w" -o /tmp/robber-datasource cmd/robber-datasource.go

# Run
FROM alpine:3.12
# Copy binary file
COPY --from=builder /tmp/robber-datasource /usr/bin/
COPY --from=builder /workspace/liarsa.com/robber-datasource/config.toml /etc/robber-datasource/config.toml

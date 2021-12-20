# Build the manager binary
FROM golang:1.17.3-buster AS builder

LABEL maintainer="eviltomorrow@163.com"

ENV WORKSPACE=/app GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY="https://goproxy.io,direct"

WORKDIR $WORKSPACE

ADD . .

# Build
RUN go build -ldflags "-X main.GitSha=${GITSHA} -X main.GitTag=${GITTAG} -X main.GitBranch=${GITBRANCH} -X main.BuildTime=${BUILDTIME} -s -w" -gcflags "all=-trimpath=${GOPATH}" -o bin/robber-datasource cmd/robber-datasource.go

# Run
FROM alpine:3.15
# Copy binary file
COPY --from=builder /app/bin/robber-datasource /bin/
COPY --from=builder /app/config/config-docker.toml /etc/robber-datasource/config-docker.toml

VOLUME ["/var/log/robber-datasource"]

EXPOSE 27320 2379 27017
ENTRYPOINT ["/bin/robber-datasource", "-c", "/etc/robber-datasource/config-docker.toml"]
FROM golang:alpine3.18 AS build

# Important:
#   Because this is a CGO enabled package, you are required to set it as 1.
ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /workspace

COPY . /workspace/

RUN \
    go mod tidy && \
    go install -ldflags='-extldflags "-static"' ./main.go
# go install -ldflags='-s -w -extldflags "-static"' ./main.go

FROM scratch

COPY --from=build /go/bin/main /usr/local/bin/main

ENV PORT=8080

WORKDIR /app

ENTRYPOINT [ "/usr/local/bin/main" ]
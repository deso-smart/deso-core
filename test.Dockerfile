FROM golang:1.17-alpine3.15 AS builder

RUN apk --no-cache add gcc g++ vips-dev upx

WORKDIR /usr/src/deso/core

COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY cmd cmd
COPY desohash desohash
COPY lib lib
COPY migrate migrate
COPY test_data test_data
COPY main.go .

RUN GOOS=linux go build -ldflags "-s -w" -o /usr/local/bin/deso-core main.go
RUN upx /usr/local/bin/deso-core

ENTRYPOINT ["go", "test", "-v", "github.com/deso-smart/deso-core/v2/lib"]

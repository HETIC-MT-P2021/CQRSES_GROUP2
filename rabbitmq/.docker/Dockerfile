FROM golang:1.14.1-alpine AS builder
LABEL stage=builder

WORKDIR /go/src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build \
    -a \
    -trimpath \
    -ldflags '-s -w -extldflags "-static"' \
    -tags 'netgo osusergo static_build' \
    -o /go/bin/app

FROM alpine
COPY --from=builder /go/bin/app .
ENTRYPOINT ["./app"]
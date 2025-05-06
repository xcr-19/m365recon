FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git build-base gcc musl-dev

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build .

FROM alpine:3.21.3
RUN apk upgrade --no-cache \
    && apk add --no-cache ca-certificates
COPY --from=builder /app/m365recon /usr/local/bin/

ENTRYPOINT ["m365recon"]

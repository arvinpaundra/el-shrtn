FROM golang:1.20.7-alpine AS builder

ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main

FROM alpine:latest

RUN apk add --no-cache tzdata \
  ca-certificates

ARG APP_VERSION

ENV APP_VERSION $APP_VERSION
ENV TZ Asia/Jakarta

WORKDIR /app

COPY --from=builder /app/main .

RUN echo "$APP_VERSION" > version.txt

CMD ["./main", "rest"]

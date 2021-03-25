# Compile app binary
FROM golang:latest as build-env

WORKDIR /go/src/github.com/jeroldleslie/go-coding-challenge-trivia
ENV GO111MODULE=on

COPY go.mod go.mod
COPY go.sum go.sum

WORKDIR /go/src/github.com/jeroldleslie/go-coding-challenge-trivia
COPY cmd cmd
COPY internal internal

ARG STACK
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -installsuffix cgo ./cmd/${STACK}/...

# Run app in scratch
FROM alpine:latest
ARG APP_NAME
COPY --from=build-env /go/bin/${APP_NAME} /app

EXPOSE 8000
CMD ["/app"]

# Build stage
FROM golang:1.8-alpine AS build-env

ENV GOPATH "/go"

ADD . /go/src/github-bot
RUN cd /go/src/github-bot && \
  go build -o github-bot

# Run stage
FROM alpine
MAINTAINER Kyle Bai(kyle.b@inwinstack.com)

COPY --from=build-env /go/src/github-bot/github-bot /bin/github-bot
RUN apk add --no-cache git openssh

ENTRYPOINT ["/bin/github-bot"]

FROM golang:1.16.4
MAINTAINER Lisa

ENV GGO111MODULE on

WORKDIR /app

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin && \
	air init

CMD go install && \
    air

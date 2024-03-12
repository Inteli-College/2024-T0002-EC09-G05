FROM golang:alpine

RUN go install github.com/cosmtrek/air@latest

ADD --chown=simulator ./certs /app/certs

WORKDIR /app

ENTRYPOINT ["air"]
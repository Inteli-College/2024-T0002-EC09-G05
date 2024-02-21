FROM golang:latest

RUN go install github.com/cosmtrek/air@latest

ADD --chown=simulator ./certs /app/certs

WORKDIR /app

ENTRYPOINT ["air"]
FROM golang:latest


# RUN cd /app && go mod download
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app


ENTRYPOINT ["air"]
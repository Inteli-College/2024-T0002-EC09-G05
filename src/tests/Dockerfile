FROM golang:latest

COPY . /app

WORKDIR /app

RUN go mod tidy

CMD ["tail", "-f", "/dev/null"]

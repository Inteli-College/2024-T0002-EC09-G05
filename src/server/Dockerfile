FROM golang:alpine


ENV GIN_MODE=debug

WORKDIR /go/src/g5.smartcities

COPY . /go/src/g5.smartcities/

RUN apk update && apk add --no-cache git
RUN go get ./...

RUN go build /go/src/g5.smartcities/test.go

EXPOSE $PORT

ENTRYPOINT ["./test"]
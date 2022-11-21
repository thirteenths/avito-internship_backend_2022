# syntax=docker/dockerfile:1

FROM golang:1.16-alpine 

WORKDIR /go/src/github.com/thirteenths/final/

COPY go.mod /go/src/github.com/thirteenths/final/
COPY go.sum /go/src/github.com/thirteenths/final/
RUN go mod download

COPY . /go/src/github.com/thirteenths/final/

RUN go build -o /go/src/github.com/thirteenths/final/docker-gs-ping

EXPOSE 8080 8080

CMD [ "/go/src/github.com/thirteenths/final/docker-gs-ping" ]

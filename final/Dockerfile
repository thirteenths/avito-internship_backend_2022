FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /go/src//
WORKDIR /go/src/github.com/thirteenths/final
RUN go mod download
COPY . /go/src/github.com/thirteenths/final
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/final github.com/thirteenths/final


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/thirteenths/final/build/final /usr/bin/final

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/final"]

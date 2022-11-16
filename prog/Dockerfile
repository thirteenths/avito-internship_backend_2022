FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/thirteenths/avito-internship_backend_2022/prog/
WORKDIR /go/src/github.com/thirteenths/avito-internship_backend_2022/prog
RUN go mod download
COPY . /go/src/github.com/thirteenths/avito-internship_backend_2022/prog
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/prog github.com/thirteenths/avito-internship_backend_2022/prog

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/thirteenths/avito-internship_backend_2022/prog/build/prog /usr/bin/prog
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/prog"]
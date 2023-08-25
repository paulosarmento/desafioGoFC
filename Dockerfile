FROM golang:1.21.0

WORKDIR /go/src/app
COPY . .

RUN go build -o /go/bin/app

CMD ["/go/bin/app"]

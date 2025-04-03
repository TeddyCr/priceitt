FROM golang:1.23-alpine
COPY ./service/go.mod ./service/go.sum /go/src/github.com/TeddyCr/priceitt/service/

WORKDIR /go/src/github.com/TeddyCr/priceitt/service/bootstrap
RUN go mod download
COPY ./service/ /go/src/github.com/TeddyCr/priceitt/service/

WORKDIR /go/src/github.com/TeddyCr/priceitt/service/bootstrap/cmd
CMD ["go", "run", "migrate.go"]

FROM golang:1.22-alpine
COPY ./go.mod ./go.sum /go/src/github.com/TeddyCr/priceitt/service/

WORKDIR /go/src/github.com/TeddyCr/priceitt/service/bootstrap
RUN go mod download
COPY . /go/src/github.com/TeddyCr/priceitt/service/

WORKDIR /go/src/github.com/TeddyCr/priceitt/service/bootstrap/cmd
CMD ["go", "run", "migrate.go"]

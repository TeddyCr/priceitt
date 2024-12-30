FROM golang:1.22-alpine
COPY ./go.mod ./go.sum /go/src/github.com/TeddyCr/priceitt/edgeAuthorizationServer/

WORKDIR /go/src/github.com/TeddyCr/priceitt/edgeAuthorizationServer/bootstrap
RUN go mod download
COPY . /go/src/github.com/TeddyCr/priceitt/edgeAuthorizationServer/

WORKDIR /go/src/github.com/TeddyCr/priceitt/edgeAuthorizationServer/bootstrap/cmd
CMD ["go", "run", "migrate.go"]

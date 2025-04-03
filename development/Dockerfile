FROM golang:1.23-alpine AS builder
RUN pwd
COPY ./service /go/src/github.com/TeddyCr/priceitt/service/

# Install dependency from local modules
WORKDIR /go/src/github.com/TeddyCr/priceitt/

WORKDIR /go/src/github.com/TeddyCr/priceitt/service/
RUN go build -o /go/bin/github.com/TeddyCr/priceitt/service
EXPOSE 8000 8000

FROM alpine:3.20
COPY --from=builder /go/bin/github.com/TeddyCr/priceitt/service /usr/bin/github.com/TeddyCr/priceitt/service
COPY --from=builder /go/src/github.com/TeddyCr/priceitt/service/bootstrap/config/config.yaml /usr/bin/config.yaml
WORKDIR /usr/bin
ENTRYPOINT ["/usr/bin/github.com/TeddyCr/priceitt/service"]
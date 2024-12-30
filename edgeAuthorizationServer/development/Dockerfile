FROM golang:1.22-alpine AS builder
RUN pwd
COPY ./edgeAuthorizationServer /go/src/github.com/TeddyCr/priceitt/edgeAuthorizationServer/
COPY ./models /go/src/github.com/TeddyCr/priceitt/models/
COPY ./utils /go/src/github.com/TeddyCr/priceitt/utils/

# Install dependency from local modules
WORKDIR /go/src/github.com/TeddyCr/priceitt/
RUN go work init ./edgeAuthorizationServer
RUN go work use ./utils
RUN go work use ./models

WORKDIR /go/src/github.com/TeddyCr/priceitt/edgeAuthorizationServer/
RUN go build -o /go/bin/github.com/TeddyCr/priceitt/edgeAuthorizationServer
EXPOSE 8000 8000

FROM alpine:3.20
COPY --from=builder /go/bin/github.com/TeddyCr/priceitt/edgeAuthorizationServer /usr/bin/github.com/TeddyCr/priceitt/edgeAuthorizationServer
COPY --from=builder /go/src/github.com/TeddyCr/priceitt/edgeAuthorizationServer/bootstrap/config/config.yaml /usr/bin/config.yaml
WORKDIR /usr/bin
ENTRYPOINT ["/usr/bin/github.com/TeddyCr/priceitt/edgeAuthorizationServer"]
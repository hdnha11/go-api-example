FROM golang:1.10.3-alpine3.7

RUN apk add --update git \
    && rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/hdnha11/go-api-example
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-api-example"]

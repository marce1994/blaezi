# Build Image
FROM golang:alpine

LABEL maintainer="pending"

RUN mkdir -p /go/src/github.com//burntcarrot/blaezi
COPY . /go/src/github.com/burntcarrot/blaezi

WORKDIR /go/src/github.com/burntcarrot/blaezi

RUN go build
RUN go get /go/src/github.com/burntcarrot/blaezi

# Execute Image
FROM alpine:3.6

RUN apk add --no-cache ca-certificates
COPY example.json ./test.json
COPY --from=0 /go/bin/blaezi  .

ENTRYPOINT ["./blaezi"]
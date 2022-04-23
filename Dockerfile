FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
RUN apk update
RUN apk upgrade
RUN apk add build-base
RUN apk add git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -a -v ./cmd/ipaddress
WORKDIR /dist
RUN cp -R /build/* .
RUN cd ./cmd/ipaddress
CMD ["./ipaddress"]

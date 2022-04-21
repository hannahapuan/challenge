FROM golang:1.18.1

RUN mkdir /challenge
ADD . /challenge
WORKDIR /challenge
RUN go mod download
RUN go build -o main .

CMD ["/challenge/main"]
FROM --platform=linux/arm64 golang:1.18

RUN mkdir /challenge
ADD . /challenge
WORKDIR /challenge
RUN go mod download
RUN go mod tidy
RUN go build -o main .
RUN chmod +x main

CMD ["/challenge/main"]


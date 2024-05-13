FROM golang:1.22.2-alpine

WORKDIR /app

ADD go.mod ./
ADD go.sum ./
ADD aws_config.ini ./config.ini

RUN go mod download

ADD ./ ./

RUN go build -o /main

CMD ["/main"]
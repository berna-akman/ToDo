FROM golang:1.20.2-alpine3.17 as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
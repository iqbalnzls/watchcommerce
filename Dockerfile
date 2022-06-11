# Builder
FROM golang:1.14.2-alpine3.11 as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o watchcommerce

CMD ["/app/watchcommerce"]
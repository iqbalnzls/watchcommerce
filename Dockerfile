# Builder
FROM golang:1.22-alpine3.19 as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o watchcommerce

CMD ["/app/watchcommerce"]
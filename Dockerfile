###################  Dev ###################
FROM golang:1.18-alpine as dev
WORKDIR /app/go-pay

RUN apk add --no-cache autoconf git build-base

COPY . .
RUN go mod download

CMD ["go", "run", "main.go"]


################ Production ################
FROM golang:1.18-alpine as prod
WORKDIR /app/go-pay

RUN apk add --no-cache git build-base

COPY . .

## download deps and build golang bin
RUN go mod download
RUN go build -o bin main.go

## run binary
CMD ["./bin/main"]

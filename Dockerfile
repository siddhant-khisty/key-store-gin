FROM golang:1.21.3-alpine as build

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go mod download

RUN go build -o gin-server

##############

FROM alpine:3.12

COPY --from=build /app /app

WORKDIR /app

CMD [ "gin-server" ]


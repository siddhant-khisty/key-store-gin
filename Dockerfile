FROM golang:1.21.3-alpine as build

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go mod download

RUN go build -o gin-server

##############

FROM alpine:3.20.1

COPY --from=build /app /app

WORKDIR /app

EXPOSE 9000

CMD [ "/app/gin-server" ]


#--------------------BUILD---------------------------------------
FROM golang:1.10 as builder

WORKDIR /go/src/github.com/tttmaximttt/microservicesEmplStarter/consignment-service

COPY . .
RUN ls ./proto -la
RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build .
#-----------------------------------------------------------------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

RUN ls -la
COPY --from=builder  /go/src/github.com/tttmaximttt/microservicesEmplStarter/consignment-service/consignment-service .

CMD ["./consignment-service"]
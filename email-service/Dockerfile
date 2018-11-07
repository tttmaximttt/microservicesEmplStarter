#--------------------BUILD---------------------------------------
FROM golang:1.10 as builder

WORKDIR /go/src/github.com/tttmaximttt/microservicesEmplStarter/email-service

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/tttmaximttt/microservicesEmplStarter/email-service/email-service .

CMD ["./email-service"]
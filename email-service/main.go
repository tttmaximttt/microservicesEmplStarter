package main

import (
  "log"

  k8s "github.com/micro/kubernetes/go/micro"
  pb "github.com/tttmaximttt/microservicesEmplStarter/user-service/proto/user"
  "github.com/micro/go-micro"
  "golang.org/x/net/context"
  _ "github.com/micro/go-plugins/broker/nats"
)

const topic = "user.created"

type Subscriber struct{}


func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
  log.Println("Picked up a new message")
  log.Println("Sending email to:", user.Name)
  return nil
}

func main() {
  srv := k8s.NewService(
    micro.Name("micros.go.micro.srv.email"),
  )

  srv.Init()

  micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

  // Run the server
  if err := srv.Run(); err != nil {
    log.Println(err)
  }
}
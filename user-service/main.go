package main

import (
  "fmt"
  "log"
  pb "github.com/tttmaximttt/microservicesEmplStarter/user-service/proto/user"
  "github.com/micro/go-micro"
)

func main() {
  db, err := CreateConnection()
  defer db.Close()

  if err != nil {
    log.Fatalf("Could not connect to DB: %v", err)
  }

  db.AutoMigrate(&pb.User{})
  repo := &UserRepository{db}
  tokenService := &TokenService{repo}

  srv := micro.NewService(

    // This name must match the package name given in your protobuf definition
    micro.Name("micros.go.micro.srv.user"),
  )

  srv.Init()
  publisher := micro.NewPublisher("user.created", srv.Client())

  pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService, publisher})

  if err := srv.Run(); err != nil {
    fmt.Println(err)
  }
}
package main

import (
  "fmt"
  "log"
  "errors"

  "golang.org/x/net/context"

  k8s "github.com/micro/kubernetes/go/micro"
  pb "github.com/tttmaximttt/microservicesEmplStarter/consignment-service/proto/consignment"
  vesselProto "github.com/tttmaximttt/microservicesEmplStarter/vessel-service/proto/vessel"
  userService "github.com/tttmaximttt/microservicesEmplStarter/user-service/proto/user"
  "github.com/micro/go-micro"
  "github.com/micro/go-micro/metadata"
  "os"
  "github.com/micro/go-micro/server"
)

const (
  defaultHost = "localhost:27017"
)

var (
  srv micro.Service
)

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
  return func(ctx context.Context, req server.Request, resp interface{}) error {

    if os.Getenv("DISABLE_AUTH") == "true" {
      return fn(ctx, req, resp)
    }

    meta, ok := metadata.FromContext(ctx)
    if !ok {
      return errors.New("no auth meta-data found in request")
    }

    token := meta["Token"]
    log.Println("Authenticating with token: ", token)

    // Auth here
    authClient := userService.NewUserServiceClient("go.micro.srv.user", srv.Client())
    _, err := authClient.ValidateToken(context.Background(), &userService.Token{
      Token: token,
    })

    if err != nil {
      return err
    }

    err = fn(ctx, req, resp)
    return err
  }
}

func main() {
  host := os.Getenv("DB_HOST")

  if host == "" {
    host = defaultHost
  }

  session, err := CreateSession(host)

  defer session.Close()

  if err != nil {
    log.Panicf("Could not connect to datastore with host %s - %v", host, err)
  }

  srv = k8s.NewService(
    micro.Name("micros.consignment"),
    micro.Version("latest"),
    micro.WrapHandler(AuthWrapper),
  )

  vesselClient := vesselProto.NewVesselServiceClient("micros.vessel", srv.Client())
  pongResponse, err := vesselClient.Ping(context.TODO(), &vesselProto.PingRequest{Ping: "ping"})

  if pongResponse == nil {
    log.Fatal("Service unavaliable micros.vessel")
    panic(err)
  }
  // Init will parse the command line flags.
  srv.Init()

  // Register handler
  pb.RegisterConsignmentServiceHandler(srv.Server(), &handler{vesselClient, session })

  // Run the server
  if err := srv.Run(); err != nil {
    fmt.Println(err)
  }

}
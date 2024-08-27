package main

import (
	"bidi-streaming/pb"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)


type userService struct {
    pb.UnimplementedUsersServer
}

func (s *userService) GetHelp(stream pb.Users_GetHelpServer) error {
    log.Println("Client connected")
    for {
        request, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }

        fmt.Printf("Request received: %s\n", request.Request)
        response := pb.UserHelpReply{
            Response: request.Request,
        }
        err = stream.Send(&response)
        if err != nil {
            return err
        }

    }
    log.Panicln("Client disconnected")
    return nil
}

func registerServices(s *grpc.Server) {
    pb.RegisterUsersServer(s, &userService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
    return s.Serve(l)
}

func main() {
    listenAddr := os.Getenv("LISTEN_ADDR")
    if len(listenAddr) == 0 {
        listenAddr = ":5000"
    }

    lis, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatal(err)
    }

    s := grpc.NewServer()
    registerServices(s)
    log.Fatal(startServer(s, lis))

}

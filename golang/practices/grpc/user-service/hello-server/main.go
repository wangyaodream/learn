package main

import (
	"context"
	"log"
	"net"
	"user-service/hello-server/pb"

	"google.golang.org/grpc"
)

type Server struct {
    pb.UnimplementedSayHelloServer 
}

func (s *Server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{ResponseMsg: "ok " + "hello " + r.RequestName}, nil
}


func main() {
    listen, err := net.Listen("tcp", ":5000")
    if err != nil {
        log.Printf("error in listen: %v", err)
    }
    
    grpcServer := grpc.NewServer()
    // register server
    pb.RegisterSayHelloServer(grpcServer, &Server{})

    // start server
    err = grpcServer.Serve(listen)
    if err != nil {
        log.Printf("error in start serve: %v", err)
    }
}

package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"strings"
	"user-service/service"

	"google.golang.org/grpc"
)


type userService struct {
    service.UnimplementedUsersServer
}

// GetUser是在proto中定义的函数，经过自动生成后成为一个必须实现的接口方法
func (s *userService) GetUser(ctx context.Context, in *service.UserGetRequest) (*service.UserGetReply, error) {
    log.Printf("Received request for user with Email: %s Id: %s\n", in.Email, in.Id)
    components := strings.Split(in.Email, "@")
    if len(components) != 2 {
        return nil, errors.New("invalid email address")
    }

    u := service.User{
        Id: in.Id,
        FirstName: components[0],
        LastName: components[1],
        Age: 30,
    }
    return &service.UserGetReply{User: &u}, nil
}

// 注册服务处理程序
func registerServices(s *grpc.Server) {
    service.RegisterUsersServer(s, &userService{})
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





package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"test-client/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 建立和服务器的连接
func setupGrpcConnection(addr string) (*grpc.ClientConn, error) {
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
    return grpc.NewClient(
        addr,
        opts...
    )
}

func getUserServiceClient(conn *grpc.ClientConn) service.UsersClient {
    return service.NewUsersClient(conn)
}

func getUser(client service.UsersClient, u *service.UserGetRequest) (*service.UserGetReply, error) {
    return client.GetUser(context.Background(), u)

}

func main() {

    if len(os.Args) != 2 {
        log.Fatal("Must specify a gRPC server address")
    }
    conn, err := setupGrpcConnection(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    defer conn.Close()

    c := getUserServiceClient(conn)

    result, err := getUser(c, &service.UserGetRequest{Email: "jane@doe.com", Id: "996"})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(os.Stdout, "User: %s %s\n", result.User.FirstName, result.User.LastName)
}

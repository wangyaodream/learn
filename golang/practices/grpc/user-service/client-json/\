package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"user-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
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

func createUserRequest(jsonQuery string) (*service.UserGetRequest, error) {
    u := service.UserGetRequest{}
    input := []byte(jsonQuery)
    return &u, protojson.Unmarshal(input, &u)
}

func getUserResponseJson(result *service.UserGetReply) ([]byte, error) {
    return protojson.Marshal(result)
}

func main() {

    if len(os.Args) != 3 {
        log.Fatal("Must specify a gRPC server address and search query")
    }
    serverAddr := os.Args[1]
    u, err := createUserRequest(os.Args[2])
    if err != nil {
        log.Fatalf("Bad user input: %v", err)
    }

    conn, err := setupGrpcConnection(serverAddr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    client := getUserServiceClient(conn)

    result, err := getUser(
        client,
        u,
    )

    if err != nil {
        log.Fatal(err)
    }

    // 将结果反序列化成json字符串
    data, err := getUserResponseJson(result)

    fmt.Fprintf(os.Stdout, "User: %s %s\n", result.User.FirstName, result.User.LastName)

    fmt.Fprintln(os.Stdout, result)
    fmt.Fprint(os.Stdout, string(data))
}

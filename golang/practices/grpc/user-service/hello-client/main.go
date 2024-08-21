package main

import (
	"context"
	"fmt"
	"log"
	"user-service/hello-client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
    conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("connection error: %v", err)
    }

    defer conn.Close()

    // create client
    client := pb.NewSayHelloClient(conn)

    resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "yaoyao"})
    if err != nil {
        log.Printf("error in get response: %v", err)
    }

    fmt.Println(resp.GetResponseMsg())

}

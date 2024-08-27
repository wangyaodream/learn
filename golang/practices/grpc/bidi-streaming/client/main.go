package main

import (
	"bidi-streaming/pb"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func setupGrpcConn(addr string) (*grpc.ClientConn, error) {
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
    conn, err := grpc.NewClient(addr, opts...)
    if err != nil {
        log.Printf("Error on connection: %v", err)
    }

    defer conn.Close()
    
    return conn, nil
}

func getUSerServiceClient(conn *grpc.ClientConn) pb.UsersClient {
    return pb.NewUsersClient(conn)
}

func setupChat(r io.Reader, w io.Writer, c pb.UsersClient) error {
    stream, err := c.GetHelp(context.Background())
    if err != nil {
        return err
    }

    for {
        scanner := bufio.NewScanner(r)
        prompt := "Request: "
        fmt.Fprint(w, prompt)

        scanner.Scan()
        if err := scanner.Err();err != nil {
            return err
        }

        msg := scanner.Text()
        if msg == "quit" {
            break
        }

        request := pb.UserHelpRequest{
            Request: msg,
        }
        err := stream.Send(&request)
        if err != nil {
            return err
        }

        resp, err := stream.Recv()
        if err != nil {
            return err
        }

        fmt.Printf("Response: %s\n", resp.Response)
    }

    return stream.CloseSend()
}

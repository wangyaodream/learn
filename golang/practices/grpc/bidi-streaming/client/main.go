package main

import (
	"bidi-streaming/pb"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

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

func main() {
    if len(os.Args) != 2 {
        log.Fatal("Must specify a gRPC server address")
    }

    conn, err := setupGrpcConn(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    c := getUSerServiceClient(conn)
    err = setupChat(os.Stdin, os.Stdout, c)
    if err != nil {
        log.Fatal(err)
    }
}

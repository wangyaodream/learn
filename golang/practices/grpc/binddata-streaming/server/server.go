package main

import (
	"binddata-streaming/pb"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type repoService struct {
    pb.UnimplementedRepoServer
}

// 实现CreateRepo方法
func (s *repoService) CreateRepo(stream pb.Repo_CreateRepoServer) error {
    var repoContext *pb.RepoContext
    var data []byte

    for {
        r, err := stream.Recv()
        if err == io.EOF {
            break
        }

        switch t := r.Body.(type) {
        case *pb.RepoCreateRequest_Context:
            repoContext = r.GetContext()
        case *pb.RepoCreateRequest_Data:
            b := r.GetData()
            data = append(data, b...)
        case nil:
            return status.Errorf(
                codes.InvalidArgument,
                "Message doesn't contain contedxt or data",
            )
        default:
            return status.Errorf(
                codes.FailedPrecondition,
                "Unexpected message type: %s",
                t,
            )
        }
    }

    // 创建响应消息
    repo := pb.Repository{
        Name: repoContext.Name,
        Url: fmt.Sprintf("https://git.example.com/%s/%s", repoContext.CreatorId, repoContext.Name),
    }
    r := pb.RepoCreateReply{
        Repo: &repo,
        Size: int32(len(data)),
    }
    return stream.SendAndClose(&r)
}

func registerServices(s *grpc.Server) {
    pb.RegisterRepoServer(s, &repoService{})
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

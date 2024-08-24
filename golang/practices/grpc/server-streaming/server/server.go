package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"server-streaming/pb"
	"strings"

	"google.golang.org/grpc"
)

type userService struct {
	pb.UnimplementedUsersServer
}

type repoService struct {
	pb.UnimplementedRepoServer
}

func (s *userService) GetUser(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetReply, error) {
	log.Printf("Received request for user with Email: %s Id: %s\n", in.Email, in.Id)
	components := strings.Split(in.Email, "@")
	if len(components) != 2 {
		return nil, errors.New("invalid email address")
	}

	u := pb.User{
		Id:        in.Id,
		FirstName: components[0],
		LastName:  components[1],
		Age:       30,
	}
	return &pb.UserGetReply{User: &u}, nil
}

func (s *repoService) GetRepos(in *pb.RepoGetRequest, stream pb.Repo_GetReposServer) error {
	log.Printf(
		"Received request for repo with CreatorId: %s Id: %s\n",
		in.CreatorId,
		in.Id,
	)
	repo := pb.Repository{
		Id: in.Id,
		Owner: &pb.User{
			Id:        in.CreatorId,
			FirstName: "Jane",
		},
	}

	cnt := 1
	for {
		repo.Name = fmt.Sprintf("repo-%d", cnt)
		repo.Url = fmt.Sprintf("https://git.example.com/test/%s", repo.Name)
		r := pb.RepoGetReply{
			Repo: &repo,
		}
		if err := stream.Send(&r); err != nil {
			return err
		}
		if cnt >= 5 {
			break
		}
		cnt++
	}
	return nil
}

func registerServices(s *grpc.Server) {
	pb.RegisterRepoServer(s, &repoService{})
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

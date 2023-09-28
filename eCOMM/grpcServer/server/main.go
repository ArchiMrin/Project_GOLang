package main

import (
	"context"

	"fmt"
	"log"
	"net"

	"github.com/ArchiMrin/Project_GOLang/eCOMM/grpcServer/profile"
	"google.golang.org/grpc"
)

type server struct {
	profile.UnimplementedProfileServiceServer
}

func (*server) CreateProfile(ctx context.Context, request *profile.Profile) (*profile.ProfileResponse, error) {
	return &profile.ProfileResponse{Id: "1", ErrorMessage: "", Success: request.Email + "Created with success"}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":4002")
	defer lis.Close()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	s := grpc.NewServer()
	profile.RegisterProfileServiceServer(s, &server{})
	fmt.Println("server listening on port 4002")
	if err := s.Serve(lis); err != nil {
		fmt.Println(err.Error())
	}
}

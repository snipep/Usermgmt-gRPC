package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/snipep/usermgmtGRPC/usermgmt"
	"google.golang.org/grpc"
)

var (
	port = ":50051"
)

type UserManagmentServer struct{
	pb.UnimplementedUserManagmentServer
}

func (s *UserManagmentServer) CreatNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main()  {
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()			//initialize new server
	pb.RegisterUserManagmentServer(s, &UserManagmentServer{})//register the server as a new grpc service
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}
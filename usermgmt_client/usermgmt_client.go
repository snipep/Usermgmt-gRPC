package main

import (
	"context"
	"log"
	"time"

	pb "github.com/snipep/usermgmtGRPC/usermgmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)


func main()  {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagmentClient(conn)			//creating new client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_user = make(map[string]int32)
	new_user["Vishnu"] = 21
	new_user["Zore"] = 20 
	for name, age :=range new_user{
		r, err := c.CreatNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil{
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf(`User Details:
			NAME: %s
			AGE: %d
			ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}

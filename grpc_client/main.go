package main

//client.go

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"rpc_client/pb"
)

const (
	address     = "localhost:50051"
)

func main() {
	// 初始化连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	// 客户端
	c := pb.NewOaClient(conn)

	//r, err := c.OaUserWorkExperience(context.Background(), &pb.OaUserExperienceRequest{Badge:"03100003",Type:"attendance"})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//for _,w:=range r.Data{
	//	fmt.Println(w.ExperienceType)
	//	fmt.Println(w.Detail)
	//}
	r, err := c.OaGetPreEntryPersonSalary(context.Background(), &pb.OaPreEntryPersonInfoRequest{OaId:"12345679"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)
}
package main

import (
	"context"
	"fmt"
	"grpc/gen/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := proto.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &proto.ResponseRequest{Msg: "hello world", Name: "MRB"})

	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Msg)
}

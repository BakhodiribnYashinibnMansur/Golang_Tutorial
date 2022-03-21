package main

import (
	"context"
	"grpc/gen/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type testApiSerever struct {
	proto.UnimplementedTestApiServer
}

func (server *testApiSerever) Echo(ctx context.Context, request *proto.ResponseRequest) (*proto.ResponseRequest, error) {
	return request, nil
}
func (server *testApiSerever) GetUser(ctx context.Context, request *proto.UserRequest) (response *proto.UserResponse, err error) {
	return response, nil
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterTestApiServer(grpcServer, &testApiSerever{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}

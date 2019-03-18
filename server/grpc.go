package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/laidingqing/stackbuild/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

//GrpcServer for running an HTTP server.
type GrpcServer struct {
	TLS  bool
	Port string
	Cert string
	Key  string
}

//ExecutorService ..
type ExecutorService struct{}

// List implements
func (serv *ExecutorService) List(r *proto.StreamRequest, stream proto.ExecutorService_ListServer) error {
	return nil
}

//ListenAndServe a grpc server.
func (s GrpcServer) ListenAndServe(ctx context.Context) error {
	creds, err := credentials.NewServerTLSFromFile("./server.pem", "../keys/server.key")
	lis, err := net.Listen("tcp", s.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Listen on with TLS + Token")
	serverOption := grpc.Creds(creds)
	var g *grpc.Server = grpc.NewServer(serverOption)
	defer g.Stop()

	proto.RegisterExecutorServiceServer(g, &ExecutorService{})
	reflection.Register(g)
	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

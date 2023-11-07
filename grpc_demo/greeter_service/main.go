/*
	grpc golang 服务端 demo
*/

package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc_demo/hellogrpc"
	"grpc_demo/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("%s rpc call Received: %v", utils.GetCurrentFunctionName(), in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloStreamReply(in *pb.HelloRequestList, stream pb.Greeter_SayHelloStreamReplyServer) error {
	log.Printf("%s rpc call Received: %v", utils.GetCurrentFunctionName(), in.GetMessage())
	for _, v := range in.GetMessage() {
		if err := stream.Send(&pb.HelloReply{Message: "Hello:" + v.GetName()}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

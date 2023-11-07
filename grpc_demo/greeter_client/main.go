/*
	grpc golang 客户端 demo
*/

package main

import (
	"context"
	"flag"
	"grpc_demo/greeter_client/pool"
	pb "grpc_demo/hellogrpc"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	// 单个连接创建示例
	// Set up a connection to the server.
	// conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v\n", err)
	// }
	// defer conn.Close()

	// 连接池示例
	client_pool, err := pool.GetPool(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("pool.GetPool error: %v\n", err)
	}
	conn := client_pool.Get()
	defer client_pool.Put(conn)

	// rpc调用示例
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 简单rpc
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("SayHello(_) error: %v\n", err)
	}
	log.Printf("Greeting: %s\n", r.GetMessage())

	// 服务端流式rpc
	hello_request_list := []*pb.HelloRequest{&pb.HelloRequest{Name: "Gary"}, &pb.HelloRequest{Name: "Helen"}}
	stream, err := c.SayHelloStreamReply(ctx, &pb.HelloRequestList{Message: hello_request_list})
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.SayHelloStreamReply(_) = _, %v\n", c, err)
		}
		log.Println(reply)
	}
}

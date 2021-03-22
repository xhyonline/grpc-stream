package main

import (
	"context"
	"fmt"
	"github.com/xhyonline/grpc-stream/server-stream/gen"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	connect, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer connect.Close()
	client := gen.NewServerStreamServerClient(connect)
	stream, err := client.DemoMethod(context.Background(), &gen.DemoParma{Name: "兰陵美酒郁金香"})
	if err != nil {
		panic(err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to recv: %v", err)
		}
		fmt.Println(reply.Message)
	}
	fmt.Println("本次调用结束")
}

package main

import (
	"context"
	"fmt"
	"github.com/xhyonline/grpc-stream/bothway-stream/gen"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

func main() {
	connect, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer connect.Close()
	client := gen.NewBothWayStreamServerClient(connect)
	stream, err := client.DemoMethod(context.Background())
	if err != nil {
		panic(err)
	}
	// 接收数据
	go func() {
		for {
			reply, err := stream.Recv()
			if err != nil {
				break
			}
			fmt.Println("客户端接收到的数据是:", reply.Message)
		}
	}()

	// 发送数据
	var count int
	for {
		count++
		err := stream.Send(&gen.DemoParma{Name: "兰陵美酒郁金香" + strconv.Itoa(count)})
		if err != nil {
			log.Printf("发送失败: %v", err)
			break
		}
		time.Sleep(time.Second)
		// 10 次后发送关闭
		if count == 10 {
			err := stream.CloseSend()
			if err != nil {
				break
			}
		}
	}
}

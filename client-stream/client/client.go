package main

import (
	"context"
	"fmt"
	"github.com/xhyonline/grpc-stream/client-stream/gen"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func main() {
	connect, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer connect.Close()
	client := gen.NewClientStreamServerClient(connect)
	stream, err := client.DemoMethod(context.Background())
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		stream.Send(&gen.DemoParma{Name: "兰陵美酒郁金香" + strconv.Itoa(i)})
		time.Sleep(time.Second)
	}
	// 告知服务端,已经发送完毕了,
	// stream.CloseAndRecv 读取完毕会关闭这个流的发送，这个方法返回最终结果。注意客户端只负责关闭流的发送。
	reply, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("failed to recv: %v", err)
	}
	fmt.Println("接收到数据:", reply.GetMessage())

	fmt.Println("本次调用结束")
}

package main

import (
	"fmt"
	"github.com/xhyonline/grpc-stream/client-stream/gen"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

// Server
type Server struct {
}

// DemoMethod 实现 proto 的方法
func (s *Server) DemoMethod(client gen.ClientStreamServer_DemoMethodServer) error {
	for {
		in, err := client.Recv()
		if err == io.EOF {
			// 一旦它检测的客户端关闭了流的发送，它则把最终结果发送给客户端，通过关闭这个流。
			// 流的关闭通过io.EOF这个error来区分。
			client.SendAndClose(&gen.DemoResponse{Message: "服务端响应xxxx.....data"})
			return nil
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		fmt.Println("服务端接收到一个人名:", in.Name)

	}
}

func main() {
	// 实例化一个 grpc 服务
	g := grpc.NewServer()
	s := new(Server)
	// 绑定
	gen.RegisterClientStreamServerServer(g, s)

	// grpc 监听在 8080 端口
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	// 服务启动
	err = g.Serve(l)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"github.com/xhyonline/grpc-stream/server-stream/gen"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

// Server
type Server struct {
}

// DemoMethod 实现 proto 的方法
func (s *Server) DemoMethod(in *gen.DemoParma, server gen.ServerStreamServer_DemoMethodServer) error {
	fmt.Println("客户端发送的人名参数是:", in.Name)
	for i := 0; i < 10; i++ {
		server.Send(&gen.DemoResponse{Message: "响应第" + strconv.Itoa(i) + "次"})
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	// 实例化一个 grpc 服务
	g := grpc.NewServer()
	s := new(Server)
	// 绑定
	gen.RegisterServerStreamServerServer(g, s)
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

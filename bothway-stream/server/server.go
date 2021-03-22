package main

import (
	"fmt"
	"github.com/xhyonline/grpc-stream/bothway-stream/gen"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

// Server
type Server struct {
}

// DemoMethod 实现 proto 的方法
func (s *Server) DemoMethod(server gen.BothWayStreamServer_DemoMethodServer) error {
	var count int
	// 启一个携程监听读事件
	go func() {
		for {
			p, err := server.Recv()
			if err != nil && err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("服务端接收错误", err)
				break
			}
			fmt.Println(p.Name)
		}
	}()
	// 持续写事件
	for {
		count++
		err := server.Send(&gen.DemoResponse{Message: "服务端的响应" + strconv.Itoa(count)})
		if err != nil {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	// 实例化一个 grpc 服务
	g := grpc.NewServer()
	s := new(Server)
	// 绑定
	gen.RegisterBothWayStreamServerServer(g, s)

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

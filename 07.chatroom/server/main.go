package main

import (
	"context"
	"gocamp/07.chatroom/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}

// 业务逻辑
func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	res := &proto.HelloReply{
		Message: "hello " + request.Name,
	}
	return res, nil
}

// 启动rpc的server服务
func start() {
	// 1.实例化server
	g := grpc.NewServer()
	// 2.注册逻辑到server中
	proto.RegisterGreeterServer(g, &Server{})
	// 3.启动server
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic("监听错误:" + err.Error())
	}

	err = g.Serve(lis)
	if err != nil {
		panic("启动错误:" + err.Error())
	}
}
func main() {
	start()
}

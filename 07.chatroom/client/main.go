package main

import (
	"context"
	"fmt"
	"gocamp/07.chatroom/proto"

	"google.golang.org/grpc"
)

// rpc调用
func clientRpc(body map[string]string) (res *proto.HelloReply, err error) {
	name := body["name"]
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpc := proto.NewGreeterClient(conn)
	response, err := rpc.SayHello(context.Background(), &proto.HelloRequest{Name: name})
	if err != nil {
		return nil, err
	}
	return response, nil
}

// 业务代码
func start() {
	body := make(map[string]string)
	body["name"] = "jeff"
	response, err := clientRpc(body)
	if err != nil {
		fmt.Println("rpc调用失败:", err)
		return
	}
	fmt.Println(response.Message)
}
func main() {
	start()
}

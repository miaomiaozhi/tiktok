package main

import (
	"fmt"
	"gateway/internal/router"
	"gateway/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 创建 gRPC 微服务
	srv := grpc.NewServer()

	// 注册
	service.RegisterFeedServer(srv, &router.FeedServer{})

	fmt.Println("register")
	// 启动服务
	lis, err := net.Listen("tcp", ":8899")
	fmt.Println("listening")

	if err != nil {
		log.Fatalf("Failed to listen grpc server: %v", err)
	}
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server: %v", err)
	}
}

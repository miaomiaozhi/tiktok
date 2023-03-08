package main

import (
	manager "feed/internal/manager"
	"feed/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	feedPort = ":8899"
)

func main() {
	// 创建 gRPC 微服务
	srv := grpc.NewServer()
	log.Printf("Create gRPC server success\n")

	// 注册
	service.RegisterFeedServer(srv, &manager.FeedServer{})
	log.Printf("Register Feed server success\n")

	// 启动服务
	lis, err := net.Listen("tcp", feedPort)
	log.Printf("Start Feed server success\n")

	if err != nil {
		log.Fatalf("Failed to listen grpc server: %v", err)
	}
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server: %v", err)
	}
}

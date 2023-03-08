package center

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "101.33.249.244:8899" // gRPC Feed 流 ip:port
)

func RegisterGRPCServer() *grpc.ClientConn {
	// 注册 gRPC 服务
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Register gRPC server error: %v", err)
		return nil
	}
	return conn
}

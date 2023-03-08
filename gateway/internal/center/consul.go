package center

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "101.33.249.244:8899"	// gRPC Feed 流 ip:port
)

func ResigterGRPCServer() *grpc.ClientConn {
	// 注册 gRPC 服务
	conn, err := grpc.DialContext(
		context.Background(),
		address,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect : %v", err)
		return nil
	}
	defer conn.Close()
	return conn
}

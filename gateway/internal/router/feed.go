package router

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	consul "gateway/internal/center"
	service "gateway/internal/service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const (
	address = "101.33.249.244:8080"
)

func (m *Manager) RouteFeed() {
	m.handle.GET("/douyin/feed/", m.feed)
}

type FeedServer struct {
	service.UnimplementedFeedServer
}

func (s *FeedServer) FeedVideo(ctx context.Context, req *service.TiktokFeedRequest) (*service.TiktokFeedResponse, error) {
	return &service.TiktokFeedResponse{
		StatusCode: int32(0),
		StatusMsg:  "grpc Test",
	}, nil
}

// 测试注册服务
func testGRPCServer() {
	s := grpc.NewServer()
	service.RegisterFeedServer(s, &FeedServer{})

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Feed net.Listen err %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Feed Server err %v", err)
	}
}

// 通过 gRPC 微服务返回 response
func (m *Manager) feed(c *gin.Context) {
	latestTimeStr := c.Query("latest_time")
	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if latestTime <= 0 {
		latestTime = time.Now().UnixMilli()
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Failed to connect: %v", err)
	// }
	// defer conn.Close()

	// client := pb.NewFeedClient(conn)
	// response, err := client.GetFeed(context.Background(), &pb.FeedRequest{Id: "123"})
	// if err != nil {
	// 	log.Fatalf("Failed to get feed: %v", err)
	// }
	// fmt.Printf("Feed: %v", response)

	conn := consul.RegisterGRPCServer()
	token := c.Query("token")

	client := service.NewFeedClient(conn)
	grpcResponse, err := client.FeedVideo(context.Background(), &service.TiktokFeedRequest{
		LatestTime: latestTime,
		Token:      token,
	})
	if err != nil {
		log.Fatalf("Failed to feed video: %v", err)
	}
	fmt.Printf("feed: %v", grpcResponse)

	var videoList []*service.FeedVideo
	feedResponse := service.TiktokFeedResponse{
		StatusCode: int32(200),
		StatusMsg:  grpcResponse.GetStatusMsg(),
		VideoList:  videoList,
		NextTime:   latestTime,
	}

	response := proto.Clone(&feedResponse)
	c.JSON(int(feedResponse.StatusCode), response)
}

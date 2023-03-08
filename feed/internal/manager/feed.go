package manager

import (
	"context"
	"feed/internal/service"
)

type FeedServer struct {
	service.UnimplementedFeedServer
}

func (s *FeedServer) FeedVideo(ctx context.Context, req *service.TiktokFeedRequest) (*service.TiktokFeedResponse, error) {
	return &service.TiktokFeedResponse{
		StatusMsg: "微服务内部测试",
	}, nil
}
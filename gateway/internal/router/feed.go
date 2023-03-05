package router

import (
	"net/http"
	"strconv"
	"time"

	pb "gateway/internal/service"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

const (
	localhost = "123"
)

func (m *Manager) RouteFeed() {
	m.handle.GET("/douyin/feed/", m.feed)
}

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

	var videoList []*pb.FeedVideo
	feedResponse := pb.TiktokFeedResponse {
		StatusCode: int32(200),
		StatusMsg:  "test",
		VideoList:  videoList,
		NextTime:   latestTime,
	}

	response := proto.Clone(&feedResponse)
	c.JSON(int(feedResponse.StatusCode), response)
}

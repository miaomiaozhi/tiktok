package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	StatusCode int32    `json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg,omitempty"`  // 返回状态描述
	VideoList  []*Video `json:"video_list,omitempty"`  // 视频列表
	NextTime   int64    `json:"next_time,omitempty"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type Video struct {
	Id            int64  `json:"id,omitempty"`             // 视频唯一标识
	Author        *User  `json:"author,omitempty"`         // 视频作者信息
	PlayUrl       string `json:"play_url,omitempty"`       // 视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count,omitempty"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty"`    // true-已点赞，false-未点赞
	Title         string `json:"title,omitempty"`          // 视频标题
}

type User struct {
	Id              int64  `json:"id,omitempty"`               // 用户id
	Name            string `json:"name,omitempty"`             // 用户名称
	FollowCount     int64  `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar,omitempty"`           //用户头像
	BackgroundImage string `json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `json:"signature,omitempty"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited,omitempty"`  //获赞数量
	WorkCount       int64  `json:"work_count,omitempty"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count,omitempty"`   //点赞数量
}

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

	var videoList []*Video
	feedResponse := FeedResponse{
		StatusCode: int32(200),
		StatusMsg:  "test",
		VideoList:  videoList,
		NextTime:   latestTime,
	}
	c.JSON(int(feedResponse.StatusCode), feedResponse)
}

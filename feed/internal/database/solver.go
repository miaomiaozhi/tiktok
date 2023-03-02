package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                 // 用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名称
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`            // 关注总数
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`      // 粉丝总数
	IsFollow        bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`                     // true-已关注，false-未关注
	Avatar          string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`                                          //用户头像
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`                                    //个人简介
	TotalFavorited  int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`   //获赞数量
	WorkCount       int64  `protobuf:"varint,10,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`                 //作品数量
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`     //点赞数量
}

type Video struct {
	PlayUrl       string `gorm:"column:play_url; type:"varchar(200)"`
	CoverUrl      string `gorm:"column:cover_url; type:"varchar(200)"`
	AuthorID      string `gorm:"column:author_id; type:"varchar(200)"`
	Id            int64  `gorm:"column:id; type:"bigint"`
	Title         string `gorm:"column:title; type:"varchar(200)"`
	FavoriteCount int64  `gorm:"column:favorite_count; type:"bigint"`
	CommentCount  int64  `gorm:"column:comment_count; type:"bigint"`
	IsFavorite    bool   `gorm:"column:is_favorite; type:"tinyint"`
}

type Solver struct {
	handle *gorm.DB // 需要处理的数据库
}

func (s *Solver) NewSolve(db *gorm.DB) {
	s.handle = db
}

func (s *Solver) QueryVideosListByLatestTime(n int, time time.Time) (videosList []*Video) {
	db := s.handle.Table("video_table")
	db.Where("created_at < ", time).Limit(n).Find(&videosList)
	return
}

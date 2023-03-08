package manager

import (
	"feed/internal/database/api"
	"fmt"
	"log"
	"time"
)

func (m *Manager) QueryVideosListByLatestTime(n int, time time.Time) (videosList []*api.Video) {
	db := m.handle.Table("videos")
	db.Where("created_at < ", time).Limit(n).Find(&videosList)
	return
}

func (m *Manager) QueryUserByUserId(userId int64) (user api.User) {
	db := m.handle.Table("users")
	db.Where("id = 123").Limit(1).Find(&user)
	return
}

func (m *Manager) TestQueryUser(userId int64) {
	log.Println(userId)
	user := m.QueryUserByUserId(userId)
	fmt.Println(user.Name)
}

package manager

type Manager struct {
}

// func (s *Manager) NewManager(db *gorm.DB) {
// 	s.handle = db
// }

// func (s *Manager) QueryVideosListByLatestTime(n int, time time.Time) (videosList []*service.FeedVideo) {
// 	db := s.handle.Table("video_table")
// 	db.Where("created_at < ", time).Limit(n).Find(&videosList)
// 	return
// }

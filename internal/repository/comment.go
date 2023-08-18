package repository

import (
	"douyin/sql"
	"douyin/vo"
	"time"
)

//--------------------------------- 评论------------------------------

// 创建
func Create_comment(id uint, vid int64, content string) vo.Comment {
	db := sql.GetMysqlDB()
	var a vo.Comment
	a.PublisherID = id
	a.VideoID = vid
	a.Create_date = time.Now().Format("2006-01-02 15:04:05")
	a.Content = content
	db.Create(&a)
	var user vo.Commonuser
	db.Where("id = ?", id).Find(&user)
	a.User = user
	var video vo.Video
	db.Where("id = ?", vid).Find(&video)
	video.Comment_count++
	db.Save(&video)
	return a
}

// 删除
func Delete_comment(id uint) {
	db := sql.GetMysqlDB()
	var vid int64
	db.Select("video_id").Model(vo.Comment{}).Where("id = ?", id).Find(&vid)
	db.Delete(&vo.Comment{}, id)
	var video vo.Video
	db.Where("id = ?", vid).Find(&video)
	video.Comment_count--
	db.Save(&video)
}

// 获取评论列表
func GetCommentList(vid int64) []vo.Comment {
	db := sql.GetMysqlDB()
	v := make([]vo.Comment, 0)
	db.Where("video_id = ?", vid).Find(&v)
	for i := range v {
		v[i].User = *GetUser(v[i].PublisherID)
	}
	return v
}

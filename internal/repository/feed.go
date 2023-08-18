package repository

import (
	"context"
	"douyin/sql"
	"douyin/vo"
	"fmt"
	"strconv"
	"time"
)

// 存储视频保存的本地地址
func SetVideoLocalPath(t string, f string) {
	rdb := sql.GetRedis()
	rdb.Set(context.Background(), t, f, -1)
}

// 获取视频保存的本地地址
func GetVideoLocalPath(t string) string {
	rdb := sql.GetRedis()
	s, e := rdb.Get(context.Background(), t).Result()
	if e != nil {
		fmt.Println(e)
		return ""
	}
	return s
}

// 保存视频
func SaveVideo(v vo.Video) {
	db := sql.GetMysqlDB()
	u := GetUser(v.AuthorID)
	u.Work_count++
	db.Create(&v)
	db.Save(u)
}

// 对视频类在数据库层面进行操作
func GetVideoList(id uint, t int64) ([]vo.Video, int64) {
	db := sql.GetMysqlDB()
	v := make([]vo.Video, 0)
	db.Limit(30).Where("create_time < ?", t).Last(&v)
	if len(v) == 0 {
		return nil, time.Now().Unix()
	}
	if id == 0 {
		for i := range v {
			v[i].Author = *GetUser(v[i].AuthorID)
		}
		return v, v[len(v)-1].Create_time
	}
	for i := range v {
		v[i].Is_favorite = Is_favorite(id, int(v[i].ID))
		v[i].Author = *GetUser(v[i].AuthorID)
	}
	return v, v[len(v)-1].Create_time
}

// 视频点赞或者取消点赞
func Set_favorite(id uint, vid int) {
	rdb := sql.GetRedis()
	err := rdb.SAdd(context.Background(), "like"+strconv.Itoa(int(id)), vid).Err()
	if err != nil {
		fmt.Println("redis:", err)
	}
	db := sql.GetMysqlDB()
	var video vo.Video
	db.Where("id = ?", vid).Find(&video)
	video.Favorite_count++
	db.Save(&video)
}
func Cancel_favorite(id uint, vid int) {
	rdb := sql.GetRedis()
	err := rdb.SRem(context.Background(), "like"+strconv.Itoa(int(id)), vid).Err()
	if err != nil {
		fmt.Println("redis:", err)
	}
	db := sql.GetMysqlDB()
	var video vo.Video
	db.Where("id = ?", vid).Find(&video)
	video.Favorite_count--
	db.Save(&video)
}

// 查询用户是否已经对该视频点赞
func Is_favorite(id uint, vid int) bool {
	rdb := sql.GetRedis()
	isMember, err := rdb.SIsMember(context.Background(), "like"+strconv.Itoa(int(id)), vid).Result()
	if err != nil {
		fmt.Println("redis:", err)
		return false
	}
	return isMember
}

// 获取视频点赞的列表
func GetFavorteList(id uint) []vo.Video {
	rdb := sql.GetRedis()
	vids, err := rdb.SMembers(context.Background(), "like"+strconv.Itoa(int(id))).Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db := sql.GetMysqlDB()
	v := make([]vo.Video, 0)
	db.Where("id in ?", vids).Find(&v)
	for i := range v {
		v[i].Is_favorite = true
		v[i].Author = *GetUser(v[i].AuthorID)
	}
	return v
}

// 获取我的投稿视频
func GetMyVideoList(id uint) []vo.Video {
	db := sql.GetMysqlDB()
	v := make([]vo.Video, 0)
	var u vo.Commonuser
	db.Where("author_id=?", id).Find(&v)
	db.Where("id = ?", id).Find(&u)
	for i := range v {
		v[i].Author = u
	}
	return v
}

package feed

import (
	"douyin/config"
	"douyin/internal/repository"
	"douyin/vo"
	"fmt"
	"time"
)

func GetVideoList(id uint, time int64) vo.Douyin_feed_response {
	vals, next_time := repository.GetVideoList(id, time)
	v := vo.Douyin_feed_response{
		Status_code: 0,
		Status_msg:  "获取成功",
		Video_list:  vals,
		Next_time:   next_time,
	}
	return v
}
func SaveVideo(f string, id uint, t string) {
	repository.SetVideoLocalPath(t, f)
	var v vo.Video
	v.Play_url = fmt.Sprintf("http://%s:%s/douyin/video?title=%s", config.ServerHost, config.ServerPort, t)
	v.Cover_url = fmt.Sprintf("http://%s:%s/douyin/a", config.ServerHost, config.ServerPort)
	v.Title = t
	v.Create_time = time.Now().Unix()
	v.AuthorID = id
	repository.SaveVideo(v)
}

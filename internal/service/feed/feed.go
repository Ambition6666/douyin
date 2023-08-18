package feed

import (
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
	v.Play_url = fmt.Sprintf("http://192.168.101.32:9090/video?title=%s", t)
	v.Cover_url = "http://192.168.101.32:9090/douyin/a"
	v.Title = t
	v.Create_time = time.Now().Unix()
	v.AuthorID = id
	repository.SaveVideo(v)
}

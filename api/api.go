package api

import (
	"douyin/config"
	"douyin/internal/repository"
	"douyin/internal/service/feed"
	"douyin/internal/service/login"
	"douyin/internal/service/user"
	"douyin/vo"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取视频列表
func GetVideoList(c *gin.Context) {
	t, _ := strconv.Atoi(c.Query("latest_time"))
	time := int64(t)
	token := c.Query("token")
	if token == "" {
		data := feed.GetVideoList(0, time)
		c.JSON(200, data)
	} else {
		_, i, _ := login.ParseToken(token)
		data := feed.GetVideoList(i, time)
		c.JSON(200, data)
	}
}

// 登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var data vo.Douyin_user_login_response
	data.Status_code, data.Status_msg, data.Token, data.User_id = login.Login(username, password)
	c.JSON(200, data)
}

// 注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	login.Register(username, password)
	var data vo.Douyin_user_register_response
	data.Status_code, data.Status_msg, data.Token, data.User_id = login.Login(username, password)
	c.JSON(200, data)
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	i, e := strconv.Atoi(c.Query("user_id"))
	if e != nil {
		v := vo.Douyin_user_response{
			Status_code: 1,
			Status_msg:  "传值失败",
			Commonuser:  vo.Commonuser{},
		}
		c.JSON(500, v)
		return
	}
	id := uint(i)
	val := user.GetUserInfo(id)
	fmt.Println(val)
	v := vo.Douyin_user_response{
		Status_code: 0,
		Status_msg:  "成功",
		Commonuser:  *val,
	}
	c.JSON(200, v)
}

// 投稿视频
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		c.JSON(401, vo.Publish_response{
			Status_code: 1,
			Status_msg:  "用户未登录",
		})
	}
	header, err := c.FormFile("data")
	if err != nil {
		c.JSON(400, vo.Publish_response{
			Status_code: 1,
			Status_msg:  err.Error(),
		})
		return
	}
	filename := header.Filename

	filepath := config.FileDir
	filepath = filepath + filename
	c.SaveUploadedFile(header, filepath)
	title := c.PostForm("title")
	_, i, _ := login.ParseToken(token)
	fmt.Println(i)
	feed.SaveVideo(filepath, i, title)
	c.JSON(200, vo.Publish_response{
		Status_code: 0,
		Status_msg:  "上传成功",
	})
}
func GetUserVideoList(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		v := vo.Douyin_video_response{
			Status_code: 1,
			Status_msg:  "没有登录",
			Video_list:  nil,
		}
		c.JSON(500, v)
		return
	}
	i, e := strconv.Atoi(c.Query("user_id"))
	if e != nil {
		v := vo.Douyin_video_response{
			Status_code: 1,
			Status_msg:  "传值失败",
			Video_list:  nil,
		}
		c.JSON(500, v)
		return
	}
	vals := repository.GetMyVideoList(uint(i))
	c.JSON(200, vo.Douyin_video_response{
		Status_code: 0,
		Status_msg:  "成功",
		Video_list:  vals,
	})
}

// 对视频喜爱或者取消喜爱
func Set_favorite(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(401, vo.Publish_response{
			Status_code: 1,
			Status_msg:  "用户未登录",
		})
	}
	_, id, _ := login.ParseToken(token)
	vid, _ := strconv.Atoi(c.Query("video_id"))
	action_type := c.Query("action_type")
	if action_type == "1" {
		repository.Set_favorite(id, vid)
	} else {
		repository.Cancel_favorite(id, vid)
	}
	c.JSON(200, vo.Publish_response{
		Status_code: 0,
		Status_msg:  "成功",
	})
}

// 获取用户喜欢的视频
func GetUserlike(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(token)
	if token == "" {
		c.JSON(401, vo.Douyin_userlike_response{
			Status_code: 1,
			Status_msg:  "没有登录",
			Video_list:  nil,
		})
	} else {
		_, id, _ := login.ParseToken(token)
		vals := repository.GetFavorteList(id)
		c.JSON(200, vo.Douyin_userlike_response{
			Status_code: 0,
			Status_msg:  "成功",
			Video_list:  vals,
		})
	}
}

// 发布评论
func PublishComment(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(401, vo.Douyin_comment_response{
			Status_code: 1,
			Status_msg:  "没有登录",
			Comment:     vo.Comment{},
		})
	} else {
		_, id, _ := login.ParseToken(token)
		video_id := c.Query("video_id")
		vid, _ := strconv.Atoi(video_id)
		action_type := c.Query("action_type")
		comment_text := c.Query("comment_text")
		if action_type == "1" {
			val := repository.Create_comment(id, int64(vid), comment_text)
			c.JSON(200, vo.Douyin_comment_response{
				Status_code: 0,
				Status_msg:  "创建成功",
				Comment:     val,
			})
		} else {
			comment_id, _ := strconv.Atoi(c.Query("comment_id"))
			repository.Delete_comment(uint(comment_id))
			c.JSON(200, vo.Douyin_comment_response{
				Status_code: 0,
				Status_msg:  "删除成功",
				Comment:     vo.Comment{},
			})
		}
	}

}

// 查找评论记录
func SearchComment(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(401, vo.Douyin_comment_list_response{
			Status_code:  1,
			Status_msg:   "没有登录",
			Comment_list: nil,
		})
	} else {
		video_id, _ := strconv.Atoi(c.Query("video_id"))
		vals := repository.GetCommentList(int64(video_id))
		c.JSON(200, vo.Douyin_comment_list_response{
			Status_code:  0,
			Status_msg:   "获取成功",
			Comment_list: vals,
		})
	}
}

// 关注用户或者取消用户
func Set_favorite_user(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(401, vo.Publish_response{
			Status_code: 1,
			Status_msg:  "没有登录",
		})
	} else {
		toid, _ := strconv.Atoi(c.Query("to_user_id"))
		_, id, _ := login.ParseToken(token)
		action_type := c.Query("action_type")
		if action_type == "1" {
			repository.Set_favorite_user(id, toid)
		} else {
			repository.Cancel_favorite(id, toid)
		}
		c.JSON(200, vo.Publish_response{
			Status_code: 0,
			Status_msg:  "成功",
		})
	}
}

// 获取关注信息列表
func Get_follow_list(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(401, vo.Douyin_follow_list_response{
			Status_code: 1,
			Status_msg:  "没有登录",
			User_list:   nil,
		})
	} else {
		_, id, _ := login.ParseToken(token)
		vals := repository.Get_follow_list(id)
		c.JSON(200, vo.Douyin_follow_list_response{
			Status_code: 0,
			Status_msg:  "成功",
			User_list:   vals,
		})
	}
}

// 获取粉丝信息列表
func Get_follower_list(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(401, vo.Douyin_follow_list_response{
			Status_code: 1,
			Status_msg:  "没有登录",
			User_list:   nil,
		})
	} else {
		_, id, _ := login.ParseToken(token)
		vals := repository.Get_follower_list(id)
		c.JSON(200, vo.Douyin_follow_list_response{
			Status_code: 0,
			Status_msg:  "成功",
			User_list:   vals,
		})
	}
}

// 发送消息
func SendMes(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(401, vo.Publish_response{
			Status_code: 1,
			Status_msg:  "没有登录",
		})
	} else {
		toid, _ := strconv.Atoi(c.Query("to_user_id"))
		_, id, _ := login.ParseToken(token)
		action_type := c.Query("action_type")
		content := c.Query("content")
		if action_type == "1" {
			repository.SaveMes(id, uint(toid), content)
			c.JSON(200, vo.Publish_response{
				Status_code: 0,
				Status_msg:  "成功",
			})
		}
	}
}

// 获取聊天记录
func GetChat(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(401, vo.Douyin_chat_response{
			Status_code:  1,
			Status_msg:   "没有登录",
			Message_list: nil,
		})
	} else {
		toid, _ := strconv.Atoi(c.Query("to_user_id"))
		_, id, _ := login.ParseToken(token)
		vals := repository.GetChat(id, uint(toid))
		c.JSON(200, vo.Douyin_chat_response{
			Status_code:  0,
			Status_msg:   "成功",
			Message_list: vals,
		})
	}
}

// 获取用户头像
func Giveavatur(c *gin.Context) {
	c.File("D:/douyin/images/R.jpg")
}

// 获取用户背景
func Giveback(c *gin.Context) {
	c.File("D:/douyin/images/OIP.jpg")
}

// 获取视频
func Video(c *gin.Context) {
	t := c.Query("title")
	f := repository.GetVideoLocalPath(t)
	c.File(f)
}

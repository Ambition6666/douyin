package routes

import (
	"douyin/api"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	douyin := r.Group("douyin")
	{
		//字节给的所有接口
		douyin.GET("/feed/", api.GetVideoList)
		douyin.GET("/user/", api.GetUserInfo)
		douyin.POST("/publish/action/", api.Publish)
		douyin.GET("/publish/list/", api.GetUserVideoList)
		douyin.POST("/user/login/", api.Login)
		douyin.POST("/user/register/", api.Register)
		douyin.POST("/favorite/action/", api.Set_favorite)
		douyin.GET("/favorite/list/", api.GetUserlike)
		douyin.POST("/comment/action/", api.PublishComment)
		douyin.GET("/comment/list/", api.SearchComment)
		douyin.POST("/relation/action/", api.Set_favorite_user)
		douyin.GET("/relation/follow/list/", api.Get_follow_list)
		douyin.GET("/relation/follower/list/", api.Get_follower_list)
		douyin.GET("/relation/friend/list/", api.Get_follower_list)
		douyin.POST("/message/action/", api.SendMes)
		douyin.GET("/message/chat/", api.GetChat)
		//以下接口是获取资源
		douyin.GET("/a", api.Giveavatur) //获取头像
		douyin.GET("/b", api.Giveback)   //获取背景
		douyin.GET("/video", api.Video)  //获取视频
	}
	return r
}

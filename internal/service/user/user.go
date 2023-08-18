package user

import (
	"douyin/internal/repository"
	"douyin/vo"
)

func GetUserInfo(id uint) *vo.Commonuser {
	return repository.GetUser(id)
}

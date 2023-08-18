package login

import "douyin/internal/repository"

func Login(n string, p string) (int, string, string, uint) {
	u := repository.SearchUser(n)
	if u.UID == 0 {
		return 1, "没有该用户", "", 0
	}
	if u.Pwd != p {
		return 1, "密码错误", "", 0
	}
	str, err := GetToken(Msk, u.UID)
	if err != nil {
		return 1, "获取token失败", "", 0
	}
	return 0, "登录成功", str, u.UID
}
func Register(n string, p string) (int, string, string, uint) {
	repository.Create_user(n, p)
	u := repository.SearchUser(n)
	str, err := GetToken(Msk, u.UID)
	if err != nil {
		return 1, "获取token失败", "", 0
	}
	return 0, "注册成功", str, u.UID
}

package repository

import (
	"context"
	"douyin/sql"
	"douyin/vo"
	"fmt"
	"strconv"
)

// 用户
func Create_user(n string, p string) {
	db := sql.GetMysqlDB()
	v := new(vo.User)
	v.Info.Name = n
	v.Pwd = p
	v.Info.Avatar = "http://192.168.101.32:9090/douyin/a"
	v.Info.Background_image = "http://192.168.101.32:9090/douyin/b"
	v.Info.Signature = "nb"
	v.Info.Total_favorited = "0"
	db.Create(v)
}

// 一个通过用户名,一个通过id获取
func SearchUser(n string) *vo.User {
	db := sql.GetMysqlDB()
	a := new(vo.Commonuser)
	b := new(vo.User)
	db.Where("name = ?", n).Find(a)
	db.Where("uid=?", a.ID).Find(b)
	return b
}
func GetUser(id uint) *vo.Commonuser {
	db := sql.GetMysqlDB()
	a := new(vo.Commonuser)
	db.Where("id=?", id).Find(a)
	return a
}

// 关注
func Set_favorite_user(mid uint, toid int) {
	rdb := sql.GetRedis()
	err := rdb.SAdd(context.Background(), "favorite_user"+strconv.Itoa(int(mid)), toid).Err()
	if err != nil {
		fmt.Println("redis:", err)
	}
	err = rdb.SAdd(context.Background(), "favorited_user"+strconv.Itoa(toid), mid).Err()
	if err != nil {
		fmt.Println("redis:", err)
	}
	db := sql.GetMysqlDB()
	me := GetUser(mid)
	me.Follow_count++
	db.Save(me)
	to := GetUser(uint(toid))
	to.Follower_count++
	db.Save(to)
}
func Cancel_favorite_user(mid uint, toid int) {
	rdb := sql.GetRedis()
	err := rdb.SRem(context.Background(), "favorite_user"+strconv.Itoa(int(mid)), toid).Err()
	if err != nil {
		fmt.Println("redis:", err)
	}
	err = rdb.SRem(context.Background(), "favorited_user"+strconv.Itoa(toid), mid).Err()
	if err != nil {
		fmt.Println("redis:", err)
	}
	db := sql.GetMysqlDB()
	me := GetUser(mid)
	me.Follow_count--
	db.Save(me)
	to := GetUser(uint(toid))
	to.Follower_count--
	db.Save(to)
}
func Get_follow_list(id uint) []vo.Commonuser {
	rdb := sql.GetRedis()
	vals, err := rdb.SMembers(context.Background(), "favorite_user"+strconv.Itoa(int(id))).Result()
	if err != nil {
		fmt.Println("redis:", err)
	}
	db := sql.GetMysqlDB()
	v := make([]vo.Commonuser, 0)
	db.Where("id in ?", vals).Find(&v)
	return v
}
func Get_follower_list(id uint) []vo.Commonuser {
	rdb := sql.GetRedis()
	vals, err := rdb.SMembers(context.Background(), "favorited_user"+strconv.Itoa(int(id))).Result()
	if err != nil {
		fmt.Println("redis:", err)
	}
	db := sql.GetMysqlDB()
	v := make([]vo.Commonuser, 0)
	db.Where("id in ?", vals).Find(&v)
	return v
}
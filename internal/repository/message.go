package repository

import (
	"douyin/sql"
	"douyin/vo"
	"time"
)

func SaveMes(me uint, to uint, c string) {
	db := sql.GetMysqlDB()
	v := new(vo.Message)
	v.Content = c
	v.FromID = me
	v.ToID = to
	v.Create_time = time.Now().Unix()
	db.Create(v)
}
func GetChat(id uint, toid uint) []vo.Message {
	db := sql.GetMysqlDB()
	v := make([]vo.Message, 0)
	db.Where("from_id = ? and to_id= ?", id, toid).Or("from_id = ? and to_id= ?", toid, id).First(&v)
	return v
}

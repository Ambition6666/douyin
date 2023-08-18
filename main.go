package main

import (
	"douyin/config"
	"douyin/internal/http/routes"
	"douyin/sql"
)

func init() {
	config.ConfigInit()
	sql.InitSql()
}
func main() {
	sql.RForm()
	r := routes.InitRoute()
	r.Run(":" + config.ServerPort)
}

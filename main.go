package main

import (
	"web_login/database"
	"web_login/routers"
)

func main() {
	database.InitMysql()

	router := routers.InitRouter()

	router.Run(":8081")
}
/*
ajax 实现
token需要验证正确与否 需要创建一个redis
子系统登录后把token存入url
然后通过js调用主系统8081的信息
*/
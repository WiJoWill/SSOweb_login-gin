package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web_login/model"
)

//存储数据库内五个用户数据的切片
var db_info [5] string

//显示页面
func DB_Info_Get(c * gin.Context){
	claims := c.MustGet("claims").(*UserClaims)
	if claims != nil {
		c.HTML(http.StatusOK, "db_info.html", gin.H{"title": "数据页"})
	}
}

//给前端显示5个用户数据
func DB_Info_Post(c * gin.Context){
	//检查下是否携带了token
	token, err := c.Cookie("token")
	if model.CheckToken(token) == false {
		fmt.Print(err)
		return
	}
	//通过用户id搜寻对应信息并存储
	var data []string
	for i := 0; i < 5; i++ {
		info := model.QueryUserInfoWithID(i+1)
		data = append(data, info)
	}
	//传递json数据
	c.JSON(http.StatusOK, gin.H{"code": 0,  "data" : data})
}

func RequestUserInfo(token string) string{
	fmt.Println(token)
	username := model.GetTokenValue(token)
	if username == "false"{
		return "Something Wrong. Please Check."
	}
	user_id := model.QueryUserWithUsername(username)
	user_status := model.QueryUserStatusWithUsername(username)
	//user_password := model.QueryUserPasswordWithID(user_id)
	return "username is " + username + "; user_id is " +
		strconv.Itoa(user_id) + "; user_status is " +
		strconv.Itoa(user_status) //+ "user_password is " + user_password
}

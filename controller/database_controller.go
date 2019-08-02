package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
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
		info := template.HTMLEscapeString(model.QueryUserInfoWithID(i+1))
		data = append(data, info)
	}
	//传递json数据
	c.JSON(http.StatusOK, gin.H{"code": 0, "UsersInfo": data})
}

func RequestUserInfo(token string, ip string) string{
	//验证请求者的ip地址是否与登录时一致，如果不一致则要求用户重新登录
	//这里没有加入redirect命令，可以自行添加
	if !model.CheckIPAndToken(ip, token){
		fmt.Println("IP与登录地址不符，请重新登录")
		return "Something Wrong. Please Check."
	}
	fmt.Println("ip地址为:" + ip)
	//
	fmt.Println(token)
	username := template.HTMLEscapeString(model.GetTokenValue(token))
	if username == "false"{
		return "Something Wrong. Please Check."
	}
	//对外
	user_id := model.QueryUserWithUsername(username)
	user_status := model.QueryUserStatusWithUsername(username)
	//user_password := model.QueryUserPasswordWithID(user_id)
	return "username is " + username + "; user_id is " +
		strconv.Itoa(user_id) + "; user_status is " +
		strconv.Itoa(user_status) //+ "user_password is " + user_password
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"web_login/model"
	"web_login/utility"
)

func RegisterGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK,"register.html",gin.H{"title":"注册页"})
}

//注册功能
func RegisterPost(c *gin.Context){
	//c.Header("Content-Type","text/javascript")
	//获取信息
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	//reassure := c.PostForm("reassure")
	//加了转义字符验证的获取信息
	username := template.HTMLEscapeString(c.PostForm("username"))
	password := template.HTMLEscapeString(c.PostForm("password"))
	reassure := template.HTMLEscapeString(c.PostForm("reassure"))
	fmt.Println(username, password, reassure)

	//判断该用户名是否被注册，如果已注册，返回错误
	id := model.QueryUserWithUsername(username)
	fmt.Println("id:",id)
	//如果id>0， 说明此用户名已存在于数据库内
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code":0,"message":"用户名已经存在"})
		return
	}

	//注册用户名和密码
	//密码是md5后的数据

	//这个注释掉的是安全算法的一种，因为测试的时候没使用，所以作为参考
	//password = utility.MD5(username + password)

	/* 这个注释掉的是加随机盐值的安全算法，后来加的，作为参考
	saltstring := strconv.FormatInt(rand.Int63(),10)
	salt := model.UserSalt{0, saltstring}
	_, errsalt := model.InsertUserSalt(salt)
	if errsalt != nil{
		c.JSON(http.StatusOK, gin.H{"code":0,"message":"加盐失败"})
		return
	}
	password = utility.MD5(password+saltstring)
	*/

	password = utility.MD5(password)
	fmt.Println("md5后：",password)

	user := model.User{0,username,password,0}
	_,err :=model.InsertUser(user)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"code":0,"message":"注册失败"})
	}else{
		c.JSON(http.StatusOK, gin.H{"code":1,"message":"注册成功"})
	}
}
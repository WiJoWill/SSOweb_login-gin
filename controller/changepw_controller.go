package controller
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"web_login/model"
	"web_login/utility"
)
func  ChangepwGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK,"change_password.html",gin.H{"title":"修改密码页"})
}
func ChangepwPost(c *gin.Context){
	//获取信息
	username := template.HTMLEscapeString(c.PostForm("username"))
	password := template.HTMLEscapeString(c.PostForm("password"))
	reassure := template.HTMLEscapeString(c.PostForm("reassure"))
	fmt.Println(username, password, reassure)

	//判断该用户名是否被注册，如果未被注册，返回错误
	id := model.QueryUserWithUsername(username)
	fmt.Println("id:",id)
	if id == -1 {
		c.JSON(http.StatusOK, gin.H{"code":0,"message":"用户名不存在，请检查"})
		return
	}

	//注册用户名和密码
	//密码是md5后的数据
	password = utility.MD5(password)
	fmt.Println("md5后：",password)

	user := model.User{id,username,password,2}
	_,err :=model.UpdateUser(user)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"code":0,"message":"修改密码失败"})
	}else{
		c.JSON(http.StatusOK, gin.H{"code":1,"message":"修改密码成功"})
	}
}
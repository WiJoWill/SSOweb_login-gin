package controller
//需要完成的功能
//登录成功后向token_controller发出获取token的请求
import (
	_ "errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"web_login/model"
	"web_login/utility"
)
func  LoginGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

//登录
func LoginPost(c *gin.Context) {

	//获取表单信息
	//username := c.PostForm("username")
	//password := c.PostForm("password")

	//转字符后的那种
	username:= template.HTMLEscapeString(c.PostForm("username"))
	password:= template.HTMLEscapeString(c.PostForm("password"))
	fmt.Println("username:", username, ",password:", password)

	id := model.QueryUserWithParam(username, utility.MD5(password))
	fmt.Println("id:", id)

	//转义字符后获取来源网站 这个需要吗
	//pre_url := template.HTMLEscapeString(c.Query("redirectURL"))*********
	//获取来源网站
	pre_url := c.Query("redirectURL")
	if id > 0 {
		status := model.QueryUserStatusWithUsername(username);
		fmt.Println("status", status)
		//如果用户状态为0，则说明用户是首次登录，请让用户修改密码
		if status == 0 {
			c.Redirect(http.StatusMovedPermanently, "/change_password")
		} else {
				user := model.User{id,username,password,2}

				token := generateToken(c, user)
				if token == ""{
					return
				}
				//生成一个子密钥
				//sub_token := generateSubToken(c, user)
				fmt.Println(pre_url)
				//为了安全，可以启用48行的命令，专门建立个sub_token。如果启用，请请将下面的token改为sub_token
				c.Redirect(http.StatusMovedPermanently, pre_url+"?sub_token="+token)
				//c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
			}
		}else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败，请检查用户名或密码"})
	}
}


// 生成Token
func generateToken(c *gin.Context, user model.User) string {
	j := JWT{
		[]byte("Admin"),
	}
	claims := UserClaims{
		strconv.Itoa(user.Id),
		user.Username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "Admin" ,             //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return ""
	}else{
		//将token存入Redis库中，key为token，value是用户名
		model.ConnectRedis()
		model.SetToken(token, user.Username)
		//获取对应ip你并且存储在redis里
		ip := c.ClientIP()
		model.SetTokenIP(ip,token)
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.Set("data",data)
	//这里的maxAge只设定了300秒，肯自行设定token于cookie中的过期时间
	c.SetCookie("token", token, 300, "/","127.0.0.1",false,true)
	return token
}

//给子系统提供token，不需要储存在cookie中
//这个功能可以暂时不启用，为了安全可以启用：将用户认证的token和提供用户信息的token进行分别
func generateSubToken(c *gin.Context, user model.User) string{
	j := JWT{
		[]byte("Admin"),
	}
	claims := UserClaims{
		strconv.Itoa(user.Id),
		user.Username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "Admin-For-UserInfo",            //签名的发行者, 携带特殊目的
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return "Wrong"
	}
	return token
}

// 登录结果
type LoginResult struct {
	Token string `json:"token"`
	model.User
}


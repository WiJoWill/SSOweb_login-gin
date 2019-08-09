package controller


/*要实现的功能
1、能够访问model并且让model访问redis去验证是否登录 本质上是验证token
2、将token结果信息返回给子系统
3、登录后
*/
import (
	_ "encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/fwhezfwhez/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"web_login/model"

	_ "web_login/model"
)
type JWT struct {
	SigningKey []byte
}
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "Admin"
)
type UserClaims struct {
	ID    string `json:"userId"`
	Username  string `json:"username"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取和设置signKey
func GetSignKey() string {
	return SignKey
}
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		//url := c.GetHeader("Referer")
		if token == "" {
			c.Abort()
			c.Redirect(302, "/login")
			return
		}
		log.Print("get token: ", token)
		model.ConnectRedis()
		//判断token是否伪造
		if model.CheckToken(token) == false{
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "错误token，token不存在密钥库内",
			})
			c.Abort()
			return
		}

		//判断携带token者的ip地址是否与原始一致
		user_ip := c.ClientIP();
		fmt.Println(user_ip)
		if model.CheckIPAndToken(user_ip, token) == false{
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "IP地址与登录地点不符，请重新登录",
			})
			c.Abort()
			c.Redirect(302,"/login")
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				c.Redirect(302, "/login")
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

// 生成一个Token
func (j *JWT) CreateToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString(j.SigningKey)

	return token_string, err
}

//给子系统生成token
func (j *JWT) CreateTokenSub (claims UserClaims) (string, error){
	sub_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	sub_token_string, err := sub_token.SignedString(j.SigningKey)
	return sub_token_string, err

}

// 解析Token，获取用户内容
func (j *JWT) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
// 更新Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func CheckUserToken (token string) bool{
	model.ConnectRedis()
	return model.CheckToken(token)
}

func GetUserFromToken (token string) string{
	model.ConnectRedis()
	return model.GetTokenValue(token)
}
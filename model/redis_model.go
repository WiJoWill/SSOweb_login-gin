package model

/*
	需要实现的功能
1、 GetToken 能够从redis库中返回
 */

import (
	_ "container/list"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var c redis.Conn
var err error

func ConnectRedis() {
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis error", err.Error())
	}else{
		fmt.Println("Succeed in Connecting to Redis")
	}
}


//在redis库中储存对应密钥
func SetToken(token string, username string){
	if c == nil{
		return;
	}
	c.Do("SELECT", 0)
	//存储一个 token，value是登录的用户名
	v, err0 := c.Do("SET", token, username)
	if err0 != nil {
		fmt.Println(err0)
		return
	}
	fmt.Println(v)
	settoken, err1 := c.Do("GET", token)

	if err1 != nil {
		fmt.Println(err1)
		fmt.Println("token并没有存储成功")
	}else{
		fmt.Printf("存储了token:")
		fmt.Println(settoken)
	}
}

//在redis库中储存对应密钥的ip地址
func SetTokenIP (ip string, token string){
	if c == nil{
		return;
	}
	c.Do("SELECT", 2)

	v, err0 := c.Do("SET", ip, token)
	if err0 != nil {
		fmt.Println(err0)
		return
	}

	fmt.Println(v)
	set_ip, err1 := redis.String(c.Do("GET", ip))

	if err1 != nil {
		fmt.Println(err1)
		fmt.Println("ip并没有存储成功")
	}else{
		fmt.Printf("存储了ip：")
		fmt.Println(ip)
		fmt.Println("这个ip对应的token是:")
		fmt.Println(set_ip)
	}
}

func CheckIPAndToken(ip string,token string) bool{
	if c == nil{
		return false
	}
	c.Do("SELECT", 2)
	_, err0 := c.Do("EXISTS", ip)

	if err0 != nil {
		fmt.Printf("web_login-redis_model func CheckIPAndToken 出错err0: " )
		fmt.Println(err0)
		return false
	}

	Set_token, err1 := redis.String(c.Do("GET", ip))
	if err1 != nil {
		fmt.Printf("web_login-redis_model func CheckIPAndToken 出错err1: ")
		fmt.Println(err1)
		return false
	}

	if Set_token != token{
		return false
	}
	return true
}


//检查密钥是否存在于redis库中
func CheckToken(token string) bool{
	if c == nil{
		return false
	}
	c.Do("SELECT", 0)
	username, err0 := c.Do("EXISTS", token)
	fmt.Println(username)
	if err0 != nil{
		return false
	}
	return true
}
//从redis库中获取对应密钥的用户名
func GetTokenValue (token string) string{
	ConnectRedis()
	c.Do("SELECT", 0)
	if c == nil{
		return "False"
	}
	username, err0:= redis.String(c.Do("GET", token))
	if err0 != nil{
		fmt.Println(err)
	}
	return username
}
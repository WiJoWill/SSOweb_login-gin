package model

import (
	"fmt"
	"strconv"
	"web_login/database"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除,  2是已修改密码
}

//生成新用户
func InsertUser(user User) (int64, error) {
	return database.ModifyDB("insert into users(username,password,status) values (?,?,?)",
		user.Username, user.Password, user.Status)
}
//修改用户的资料信息
func UpdateUser(user User)(int64, error){
	return database.ModifyDB("UPDATE users SET password = ?, status = ? WHERE username = ?",
		user.Password, user.Status, user.Username)
}

//按条件查询对应用户
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	if id == 0{
		return -1;
	}
	return id
}
//按条件获取对应用户的状态
func QueryUserStatusWightCon(con string) int{
	sql := fmt.Sprintf("select status from users %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	status := 0
	row.Scan(&status)
	return status
}
//根据用户名查询状态
func QueryUserStatusWithUsername(username string) int{
	sql:= fmt.Sprintf("where username='%s'", username)
	return QueryUserStatusWightCon(sql)
}
//根据用户名查询对应id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}
//根据用户名和密码，查询id
func QueryUserWithParam(username ,password string)int{
	sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
	return QueryUserWightCon(sql)
}



//根据用户id返回用户名和状态
func QueryUserInfoWithID(id int) string{
	safe_sql := fmt.Sprintf("where id = '%s'", strconv.Itoa(id))
	sql := fmt.Sprintf("select username from users %s", safe_sql)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	username := ""
	row.Scan(&username)
	status := QueryUserStatusWithUsername(username)
	return "user's id is " + strconv.Itoa(id) +
		" ; user's username is " + username +
		" ; user's status is "+ strconv.Itoa(status)
}
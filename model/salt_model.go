package model

import (
"fmt"
"strconv"
"web_login/database"
)

type UserSalt struct {
	Id     int
	Username string
	Saltstring   string
}

//生成新用户盐
func InsertUserSalt(salt UserSalt) (int64, error) {
	return database.ModifyDBSalt("insert into salt(username, salt) values (?,?)",
		salt.Username, salt.Saltstring)
}
//修改用户盐的资料信息
func UpdateUserSalt(salt UserSalt) (int64, error){
	return database.ModifyDB("UPDATE users SET username = ?, salt = ? WHERE id = ?",
		salt.Username, salt.Saltstring, salt.Id)
}

//根据用户id返回用户盐
func QueryUserSaltWithID(id int) string{
	safe_sql := fmt.Sprintf("where id = '%s'", strconv.Itoa(id))
	sql := fmt.Sprintf("select salt from salt %s", safe_sql)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	salt := ""
	row.Scan(&salt)
	return salt
}

//根据用户名返回用户盐
func QueryUserSaltWithUsername(username string) string{
	safe_sql := fmt.Sprintf("where username = '%s'", username)
	sql := fmt.Sprintf("select salt from salt %s", safe_sql)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	salt := ""
	row.Scan(&salt)
	return salt
}


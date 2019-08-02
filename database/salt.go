package database

import (
"database/sql"
"fmt"
_"github.com/go-sql-driver/mysql"
"log"
)
var saltdb *sql.DB

//用盐表的话记得在main里加
func InitMysqlforSalt() {
	fmt.Println("InitMysql....")
	if saltdb == nil {
		saltdb, _ = sql.Open("mysql", "root:will20000324@tcp(127.0.0.1:3306)/web_login?charset=utf8")
		CreateTableWithSalt()
	}
}

//创建用户盐表
func CreateTableWithSalt() {
	sql := `CREATE TABLE IF NOT EXISTS salt(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
        salt VARCHAR(64)
        );`
	ModifyDB(sql)
}
//执行数据库
func ModifyDBSalt(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
//查询数据库
func QueryRowDBSalt(sql string) *sql.Row{
	return saltdb.QueryRow(sql)
}


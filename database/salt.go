package database

import (
"database/sql"
_"fmt"
_"github.com/go-sql-driver/mysql"
"log"
)

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


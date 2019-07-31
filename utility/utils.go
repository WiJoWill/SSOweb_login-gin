package utility

import (
	"crypto/md5"
	"fmt"
)
//如果数据不一致，那MD5之后的32位会不同
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}
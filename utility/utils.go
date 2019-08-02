package utility

import (
	"crypto/md5"
	"fmt"
)
//如果数据不一致，那MD5之后的32位会不同
func MD5(str string) string {
	//加一个盐
	md5salt := fmt.Sprintf("%x", md5.Sum([]byte("pwis"+str)))
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str + md5salt)))
	return md5str
}

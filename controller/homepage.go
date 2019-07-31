package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//一个用来测试的homepage
func HomePage(c *gin.Context) {
	claims := c.MustGet("claims").(*UserClaims)
	if claims != nil {
		c.HTML(http.StatusOK, "Home.html", gin.H{"title": "首页"})
	}
}
/*
func HomePost(c * gin.Context){

}

 */

package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myusername/OceanLearn/common"
	"github.com/myusername/OceanLearn/model"
	"github.com/myusername/OceanLearn/util"

	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	passsword := c.PostForm("password")
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "手机号必须11位"})
		return
	}
	if len(passsword) < 10 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "423", "msg": "密码不能小于10位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)

	}
	log.Println(name, telephone, passsword)
	if isTelephoneExist(telephone, db) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "423", "msg": "手机号已经存在"})
		return

	}
	NewUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  passsword,
	}
	db.Create(&NewUser)
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isTelephoneExist(telephone string, db *gorm.DB) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

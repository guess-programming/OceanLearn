package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11)";not null;unique`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitMysql()

	g := gin.Default()

	g.POST("/api/auto/register", func(c *gin.Context) {
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
			name = RandomString(10)

		}
		log.Println(name, telephone, passsword)
		if isTelephoneExist(telephone, db) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "423", "msg": "手机号已经存在"})
			return

		}
		NewUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  passsword,
		}
		db.Create(&NewUser)
		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(g.Run())
}
func isTelephoneExist(telephone string, db *gorm.DB) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM._-")
	results := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range results {
		results[i] = letters[rand.Intn(len(letters))]
	}
	return string(results)
}

func InitMysql() *gorm.DB {
	var sqlName = "root"
	var password = "199610"
	var ip = "127.0.0.1"
	var port = "3306"
	var database = "ginessential"
	dataSourceName := sourceName(sqlName, password, ip, port, database)
	//fmt.Println(dataSourceName)

	conn, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印所有sql
	})
	if err != nil {
		panic(err.Error())
		log.Println(err)
		return nil
	}

	conn.AutoMigrate(&User{})
	return conn
}
func sourceName(sqlName, password, ip, port, database string) string {
	var conn string

	conn = sqlName + ":" + password + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8&parseTime=True"
	return conn

}

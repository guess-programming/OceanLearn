package common

import (
	"log"

	"github.com/myusername/OceanLearn/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
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

	conn.AutoMigrate(&model.User{})

	return conn
}
func GetDB() *gorm.DB {
	DB = InitDB()
	return DB
}
func sourceName(sqlName, password, ip, port, database string) string {
	var conn string

	conn = sqlName + ":" + password + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8&parseTime=True"
	return conn

}

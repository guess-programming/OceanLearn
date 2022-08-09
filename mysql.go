package main

import (
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	var sqlName = "root"
	var password = "199610"
	var ip = "127.0.0.1"
	var port = "3306"
	var database = "ginessential"
	dataSourceName := sourceName(sqlName, password, ip, port, database)
	//fmt.Println(dataSourceName)
	conn, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
		log.Println(err)
		return
	}

	if err = conn.Ping(); err != nil {
		log.Println("pong错误", err)
		return
	} else {
		fmt.Println("connection mysql success!")
	}
	conn.AutoMigrate(&User{})
	return conn
}
func sourceName(sqlName, password, ip, port, database string) string {
	var conn string
	conn = sqlName + ":" + password + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8&parseTime=True"
	return conn

}

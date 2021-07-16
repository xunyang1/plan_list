package dao //存放数据库相关操作

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	"plan_list/models"

	//"plan_list/models"
)

var DB *gorm.DB

func Mysql_con(){
	DB, _ = gorm.Open("mysql", "root:root1234@(192.168.99.100:13306)/plan_list?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil{
	//	fmt.Println("err is ",err)
	//}
	//创建表
	DB.AutoMigrate(&models.Demo{})//创建相应表
}

func Close(){
	DB.Close()
}
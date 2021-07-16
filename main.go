package main

import (
	//"fmt"
	//"github.com/gin-gonic/gin"
	"plan_list/routers"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"plan_list/controller"
	"plan_list/dao"
)

func main() {
	dao.Mysql_con()//连接数据库
	defer dao.Close()
	r := routers.SetupRouter()//设置路由
	r.Run()
}

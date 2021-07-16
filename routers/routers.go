package routers

import (
	"github.com/gin-gonic/gin"
	"plan_list/controller"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")//解析模板
	r.Static("/static","./static")
	r.GET("/", controller.Indexhandler)//注意引入函数没有()
	r.GET("/v1/todo", controller.Gethandler)	//查看
	r.POST("/v1/todo", controller.Posthandler) //增加
	r.PUT("/v1/todo/:id", controller.Puthandler) //更新
	r.DELETE("/v1/todo/:id", controller.Deletehandler) //删除
	return r
}
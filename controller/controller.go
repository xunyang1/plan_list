package controller  //存放路由的响应函数

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"plan_list/dao"
	"plan_list/models"
	"strconv"

	//"github.com/jinzhu/gorm"
)

func Indexhandler(c *gin.Context) {
	c.HTML(200,"index.html",nil)
}


func Gethandler(c *gin.Context) {//查询
	var users []models.Demo
	fmt.Println("-----------------------")
	dao.DB.Find(&users)//查询结果放入users
	c.JSON(200,users)
}


func Posthandler(c *gin.Context) {//创建
	fmt.Println("---------")
	var todo models.Demo
	c.ShouldBind(&todo)//参数绑定，自动生成参数对应的结构体
	err := dao.DB.Create(&todo).Error
	//post的json可返回也可以不返回因为前端没有相应的处理接口
	if err != nil{
		panic(err)
		c.JSON(200,gin.H{"err":err.Error()})
	}else {
		c.JSON(200,todo)
	}
	//body := make([]byte,1024)
	//bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	//n,_ := c.Request.Body.Read(body)
	//c.Request.Body = ioutil.NopCloser(bytes.NewReader(body[:n]))
	//if err != nil{
	//	fmt.Println("err is ",err)
	//}
	//fmt.Println(string(body[10:n-2]))
	//fmt.Println(string(bodyBytes))
	//fmt.Println("---------")
	//body := demo{
	//	ID:     1,
	//	Title:  "222",
	//	Status: true,
	//}
}


func Puthandler(c *gin.Context) {//更新
	fmt.Println("----------")
	var update models.Demo
	c.ShouldBind(&update)//参数绑定没有绑定id，所以id需要手动设置
	//fmt.Println(update)
	id := c.Param("id")
	id1, _ := strconv.Atoi(id)
	update.ID = int64(id1)
	//update := demo{
	//	ID:     int64(id1),
	//}
	//if
	dao.DB.Model(&update).UpdateColumn("status", update.Status)//无hooks更新
	//_ = c.BindJSON(&update)
	//fmt.Printf("%T",update.ID)
	//fmt.Println(update)
	//err := db.Save(&update).Error
	//if err != nil{
	//	fmt.Println("err is ",err)
	//}
	//del := demo{
	//	ID:  int64(id1),
	//}
	//err := db.Delete(&del).Error.Error()
	//if err != ""{
	//	c.JSON(200,gin.H{"err":err})
	//}else {
	//	c.JSON(200,gin.H{"err":"yes"})
	//}
}


func Deletehandler(c *gin.Context) {//删除
	id := c.Param("id")
	if id == ""{//判断id字段是否为空，若不为空才能进行删除操作，因为delet函数当id为空时删除整个表
		return
	}else {
		id1, _ := strconv.Atoi(id)
		var del models.Demo
		del.ID = int64(id1)
		dao.DB.Delete(&del)//
	}
}


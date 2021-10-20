/**
    @author: zzg
    @date: 2021/10/20 20:13
    @dir_path: routers
    @note:
**/

package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetPouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static","static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//waiting to do
		//add
		v1Group.POST("/todo", controller.AddTodo)

		//query all
		v1Group.GET("/todo", controller.QueryAllTodo)

		//modify one
		v1Group.PUT("/todo/:id", controller.ModifyTodo)

		//delete one
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}

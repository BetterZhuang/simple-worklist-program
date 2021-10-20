/**
    @author: zzg
    @date: 2021/10/20 19:41
    @dir_path: controller
    @note:
**/

package controller

import (
	"bubble/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//url -> controller -> service -> dao -> Models 【请求来了 -> 控制器 -> 业务逻辑 -> 模型层的增删改查】

// IndexHandler 返回访问页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// AddTodo 添加数据
func AddTodo(c *gin.Context) {
	//前端页面填写待办事项 点击提交 会发送请求到这里
	//1.从请求中提取出数据
	var todo Models.Todo
	c.BindJSON(&todo)

	//2.存入数据库
	err := Models.AddTodo(&todo)
	//if err != nil {
	//	return
	//}
	//3.返回响应
	if err != nil { //与上一步存入数据库合并
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK,gin.H{
		//	"code":2000,
		//	"message":"successful",
		//	"data":todo,
		//})
	}
}

// QueryAllTodo 查看所有数据
func QueryAllTodo(c *gin.Context) {
	//查看todo这个表里的所有数据

	todolist, err := Models.QueryAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

// ModifyTodo 删除待办事项
func ModifyTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	todo, err := Models.QueryByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo) //数据绑定到todo中  ★■●★■对bindjson()还有疑问
	err = Models.SaveModify(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// DeleteTodo 删除待办事项
func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id") //c.Params.Get()获取path(URL路径)参数,且获取的是string格式
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := Models.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}

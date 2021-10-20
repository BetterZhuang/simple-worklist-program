/**
    @author:zzg
    @date:2021/10/20
    @dir_path:
    @note:
**/
package main

import (
	"bubble/Models"
	"bubble/dao"
	"bubble/routers"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main()  {
	//创建数据库
	//sql：CREATE DATABASE bubble
	//connect to db
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("connect to mysql success")
	}
	defer dao.DB.Close()  //program exit -> close db


	//Model bind,自动迁移（把结构体和数据表进行对应）
	dao.DB.AutoMigrate(&Models.Todo{})


	//注册路由
	r := routers.SetPouter()
	r.Run(":10010") //启动

}

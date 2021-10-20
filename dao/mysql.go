/**
    @author: zzg
    @date: 2021/10/20 19:48
    @dir_path: dao
    @note:
**/

package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// DB 定义全局的DB对象
var (
	DB *gorm.DB
)

// InitMySQL init connection to mysql
func InitMySQL()(err error)  {
	dsn := "root:20211017@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("connect to mysql failed,err:%v\n", err)
		return
	}
	//测试是否能够连通
	return DB.DB().Ping()
}
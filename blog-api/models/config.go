package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:Canbee2018!@tcp(111.230.25.51:3306)/blog-go?charset=utf8&loc=Asia%2FShanghai")
	logs.Info("\n连接数据库成功!")
}

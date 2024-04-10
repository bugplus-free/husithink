package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Userinfo struct {
	Id       int    `orm:"auto"`
	UserName string  `orm:"unique"`//显示说明用户名为主键
	Email    string //电子邮箱
	Password string //密码
}

var DB orm.Ormer

func init() {

	// 设置默认数据库
	//数据库存放位置：./datas/husi_user.db ， 数据库别名：default
	
}

// 目前实现注册功能中的用户名是否在数据库中，因为用户名是主键唯一，如果存在，返回false,不在就插入并且返回true
func If_In_Sqlite3(user *Userinfo) bool {
	u := Userinfo{}
	qs:=DB.QueryTable("Userinfo")
	err := qs.Filter("UserName", user.UserName).One(&u)//找到了就返回非空
	if err != nil||u.Password!=user.Password {
		return false
	} else {
		fmt.Println("err--数据库中已经存在该用户")
		return true
	}
}
func If_Add_Sqlite3(user *Userinfo) bool {
	if !If_In_Sqlite3(user) {
		u:=Userinfo{
			UserName: user.UserName,
			Email: user.Email,
			Password: user.Password,
		}
		_, err := DB.Insert(&u)
		if err != nil {
			fmt.Println("err--数据库添加用户失败")
			return false
		} else {
			fmt.Println("数据库添加用户成功")
			return true
		}
	} else {
		return false
	}
}
func init(){
	orm.RegisterDataBase("default", "sqlite3", "./datas/husi_user.db")
	//注册表
	orm.RegisterModel(new(Userinfo))
	orm.SetMaxIdleConns("default", 50) //最大空闲连接
	orm.SetMaxOpenConns("default", 40) //最大数据库连接数
	orm.Debug = true
	o := orm.NewOrm()
	DB = o
	orm.RunSyncdb("default", false, true)

	
}
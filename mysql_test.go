package main

import (
	"database/sql"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestMysql(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		t.Errorf("出现错误")
	}
	checkErr(t,err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	if err != nil {
		t.Errorf("出现错误")
	}
	checkErr(t,err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(t,err)

	id, err := res.LastInsertId()
	checkErr(t,err)

	fmt.Println(t,id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(t,err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(t,err)

	affect, err := res.RowsAffected()
	checkErr(t,err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(t,err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(t,err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(t,err)

	res, err = stmt.Exec(id)
	checkErr(t,err)

	affect, err = res.RowsAffected()
	checkErr(t,err)

	fmt.Println(affect)

	db.Close()
}
func checkErr(t *testing.T,err error) {
	if err != nil {
		t.Errorf("出现错误")
	}
}
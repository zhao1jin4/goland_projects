package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//安装  go get github.com/go-sql-driver/mysql 会git clone到GOPATH下
	//导入 import (	_ "github.com/go-sql-driver/mysql")
	db, err := sql.Open("mysql", "zh:123@tcp(127.0.0.1:3306)/mydb?charset=utf8")
	if err == nil {
		fmt.Println("连接成功", db)
	} else {
		fmt.Println("连接错误", err)
	}
	defer db.Close()

	drop := "drop  table  if exists stu"
	create := "create table stu( id int ,name varchar(30))"
	execDB(db, drop)
	execDB(db, create)

	//--插入 动态参数
	stu := [2][2]string{{"3", "lisi"}, {"4", "王"}}
	stmt, _ := db.Prepare("insert into stu values (?,?)")
	for _, s := range stu {
		stmt.Exec(s[0], s[1])
	}
	//--查询 一行数据
	var id, name string
	row := db.QueryRow("select * from stu where id=4")
	row.Scan(&id, &name)
	fmt.Println(id, "--", name)

	//---查询 多行查询
	//rows, _ := db.Query("select * from stu")
	stmt, _ = db.Prepare("select * from stu where id>?")
	rows, _ := stmt.Query(-1)

	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, "--", name)
	}

}

//create ,update,delete
func execDB(db *sql.DB, sql string) {
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println("Exec执行失败", err)
		return
	}
	effectRow, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected执行失败", err)
		return
	}
	fmt.Printf("sql=%s影响行数:%d\n", sql, effectRow)
}

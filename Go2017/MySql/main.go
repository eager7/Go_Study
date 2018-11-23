package main

import (
	. "github.com/eager7/go/log"
	. "github.com/eager7/go/errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
var table =
	`CREATE TABLE userinfo (
uid INT(10) NOT NULL AUTO_INCREMENT,
username VARCHAR(64) NULL DEFAULT NULL,
departname VARCHAR(64) NULL DEFAULT NULL,
created VARCHAR(64) NULL DEFAULT NULL,
PRIMARY KEY (uid)
)CHARACTER SET = UTF8`


func main(){
	Debug.Println("sql test")
	db, err := sql.Open("mysql", "root:pct@tcp(127.0.0.1:3306)/?charset=utf8")
	CheckResult(err)
	_,err = db.Query("DROP DATABASE IF EXISTS test")
	CheckResult(err)

	_,err = db.Query("CREATE DATABASE test")
	CheckResult(err)

	db, err = sql.Open("mysql", "root:pct@tcp(127.0.0.1:3306)/test?charset=utf8")
	CheckResult(err)
	_,err = db.Exec(table)
	CheckResult(err)

	id_pan := Insert(db, "panchangtao", "微电研发", time.Now())
	Insert(db, "fanyongsheng", "微电研发", time.Now())
	Insert(db, "donghongsong", "微电研发", time.Now())
	Insert(db, "pengtao", "微电研发", time.Now())

	stmt, err := db.Prepare("update userinfo SET username=? where uid =?")
	res,err:=stmt.Exec("潘长涛", id_pan)
	affect,err:= res.RowsAffected()
	Debug.Println(affect)

	rows,err:=db.Query("SELECT * FROM userinfo")
	CheckResult(err)
	for rows.Next(){
		var uid int
		var username,departname,created string
		err:=rows.Scan(&uid, &username, &departname, &created)
		CheckResult(err)
		Info.Println(uid, username, departname, created)
	}

	stmt,err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	CheckResult(err)
	res,err = stmt.Exec(id_pan)
	CheckResult(err)
	Debug.Println(res.RowsAffected())
	db.Close()
}

func Insert(db *sql.DB, username, departname string, created time.Time)(index int64){
	stmt, err := db.Prepare("Insert userinfo Set username=?,departname=?,created=?")
	CheckResult(err)
	ret, err := stmt.Exec(username, departname, created.String())
	CheckResult(err)
	index,_ = ret.LastInsertId()
	return index
}
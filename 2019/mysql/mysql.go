package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const TableString = "`tableTest`(`id` INT NOT NULL AUTO_INCREMENT,`name` VARCHAR(100) NOT NULL,`age` INT DEFAULT 0,`birthday` DATE,PRIMARY KEY(`id`))ENGINE=InnoDB DEFAULT CHARSET=utf8;"

type DB struct {
	db *sql.DB
}

func InitializeDB() (*DB, error) {
	db, err := sql.Open("mysql", "root:zJY121123!@tcp(127.0.0.1:3306)/plainchant")
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func (db *DB) CreateTable(table string) error {
	command := "create table if not exists" + table
	ret, err := db.db.Exec(command)
	if err != nil {
		return err
	}
	fmt.Println(ret)
	return nil
}

func (db *DB) Insert() error {
	var cmd = `insert into tableTest ( name, age) values ( "pct", 31)`
	ret, err := db.db.Exec(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	return nil
}

func (db *DB) Show() {
	command := "select * from tableTest;"
	rows, err := db.db.Query(command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println(rows.Columns())
		fmt.Println(rows.Scan())
	}
}
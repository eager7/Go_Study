package main

import (
	"database/sql"
	"fmt"
)

func main() {
	dataSourceName := "root:zJY121123!@tcp(127.0.0.1:3306)/eos_park_canada_2?charset=utf8mb4&parseTime=true&loc=Local"
	MySQLInlineActionClient, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("sql.Open command failed. err:%s, dataSourceName: %s", err.Error(), dataSourceName))
	}
	
}

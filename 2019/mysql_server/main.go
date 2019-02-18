package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
)

type logger struct{}

func (logger) Print(v ...interface{}) {
	fmt.Println(v)
}
func (EOSTokenPriceInfo) TableName() string {
	return "t_token_price_info"
}

type EOSTokenPriceInfo struct {
	Index        int64
	Symbol       string
	IssueAccount string
	EOSPrice     float64
	USDPrice     float64
	Source       string
	Timestamp    string
}

func main() {
	gormDB := InitializeGorm()
	defer checkError(gormDB.Close())
	var modelTokenPrices EOSTokenPriceInfo

	if err := gormDB.Find(&modelTokenPrices).Order("`index` DESC").Error; err != nil {
		panic(err)
	}
	fmt.Println(modelTokenPrices)
}

func InitializeGorm() *gorm.DB {
	dataSourceName := "root:zJY121123!@tcp(127.0.0.1:3306)/eos_park_canada_2?charset=utf8mb4&parseTime=true&loc=Local"
	MySQLInlineActionClient, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("sql.Open command failed. err:%s, dataSourceName: %s", err.Error(), dataSourceName))
	}
	MySQLInlineActionClient.SetMaxOpenConns(2)
	MySQLInlineActionClient.SetMaxIdleConns(1)

	gormDefaultDB, err := gorm.Open("mysql", MySQLInlineActionClient)
	if err != nil {
		// gorm 会自己 ping 一次 DB
		fmt.Println("sql.Ping command failed. err:", err,
			" data_source_name: ", dataSourceName)
		panic(err)
	}
	gormDefaultDB.DB().SetMaxOpenConns(2)
	gormDefaultDB.DB().SetMaxIdleConns(1)
	gormDefaultDB.LogMode(true).SetLogger(logger{})
	return gormDefaultDB
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

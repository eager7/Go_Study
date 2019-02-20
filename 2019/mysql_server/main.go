package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
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

type EosTokenGroup struct {
	Symbol    string
	Timestamp string
}

func (EosTokenGroup) TableName() string {
	return "t_token_price_info"
}

type mysql struct {
	sql  *sql.DB
	gorm *gorm.DB
}

func main() {
	db := InitializeGorm()
	defer checkError(db.sql.Close, db.gorm.Close)

	SearchAll(db)

	if false {
		if err := SearchList(db, []TokenRequest{{
			symbol:  "SVN",
			account: "eoseventoken",
		}, {
			symbol:  "POKER",
			account: "eospokercoin",
		}}...); err != nil {
			panic(err)
		}
	}

}

func InitializeGorm() *mysql {
	dataSourceName := "root:zJY121123!@tcp(127.0.0.1:3305)/" +
		"eos_park_canada_2?charset=utf8mb4&parseTime=true&loc=Local"
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
	fmt.Println("connect sql success")
	return &mysql{
		sql:  MySQLInlineActionClient,
		gorm: gormDefaultDB,
	}
}

func SearchAll(db *mysql) {
	//var modelTokenPrices []EOSTokenPriceInfo
	var modelTokenPrices []EOSTokenPriceInfo
	start := time.Now().UnixNano()
	if err := db.gorm.Raw("select * from t_token_price_info where (symbol,issue_account,timestamp) in " +
		"(select symbol, issue_account, max(timestamp) from t_token_price_info group by symbol, issue_account)").Scan(&modelTokenPrices); err != nil {
		fmt.Println(err)
	}

	end := time.Now().UnixNano()
	fmt.Println("find result time:", (end-start)/1000000, "ms")
	fmt.Println(len(modelTokenPrices))
}

type TokenRequest struct {
	symbol  string
	account string
}

func SearchList(db *mysql, tokens ...TokenRequest) error {
	var modelTokenPrices []EOSTokenPriceInfo
	start := time.Now().UnixNano()
	for _, token := range tokens {
		fmt.Println("search token:", token)
		var modelTokenPrice EOSTokenPriceInfo
		if err := db.gorm.Where("symbol = ? AND issue_account = ?", token.symbol, token.account).
			Order("`index` DESC").Limit(1).Find(&modelTokenPrice).Error; err != nil {
			return err
		}
		modelTokenPrices = append(modelTokenPrices, modelTokenPrice)
	}
	end := time.Now().UnixNano()
	fmt.Println("find result time:", (end-start)/1000000, "ms")
	fmt.Println("result:", modelTokenPrices)
	return nil
}

func checkError(callBacks ...func() error) {
	for _, callBack := range callBacks {
		if err := callBack(); err != nil {
			fmt.Println(err)
		}
	}
}

package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
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
	Timestamp    BlockTime
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

	if true {
		if err := SearchList(db, []TokenRequest{{
			symbol:  "SVN",
			account: "eoseventoken",
		}, }...); err != nil {
			panic(err)
		}
	} else {
		SearchAll(db)
	}

}

func InitializeGorm() *mysql {
	dataSourceName := "root:zJY121123!@tcp(127.0.0.1:3306)/" +
		"eos_park?charset=utf8mb4&parseTime=true&loc=Local&parseTime=true"
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
		gorm: gormDefaultDB.LogMode(true),
	}
}

func SearchAll(db *mysql) {
	//var modelTokenPrices []EOSTokenPriceInfo
	var modelTokenPrices []EOSTokenPriceInfo
	start := time.Now().UnixNano()
	/*if err := db.gorm.Raw("select * from t_token_price_info where (symbol,exec_account,`index`) in " +
		"(select symbol, exec_account, max(`index`) from t_token_price_info group by symbol, exec_account)").Scan(&modelTokenPrices); err != nil {
		fmt.Println(err)
	}*/
	//db.gorm.Select()

	end := time.Now().UnixNano()
	fmt.Println("find result time:", (end-start)/1000000, "ms")
	fmt.Println(len(modelTokenPrices))
	//fmt.Println(modelTokenPrices)
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
		if err := db.gorm.Where("symbol = ? AND exec_account = ?", token.symbol, token.account).
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

type BlockTime struct {
	time.Time
}
const BLOCK_TIME_FORMAT = "2006-01-02T15:04:05.000"

func NewBlockTime(t time.Time) BlockTime {
	return BlockTime{Time: t}
}

func (this BlockTime) MarshalJSON() ([]byte, error) {
	t := this.UTC().Format(BLOCK_TIME_FORMAT)
	var stamp = fmt.Sprintf("\"%s\"", t)
	return []byte(stamp), nil
}

func (this *BlockTime) UnmarshalJSON(b []byte) error {
	// timeStr := string(b)
	timeStr := strings.Replace(string(b), "\"", "", -1)
	d, err := time.Parse(BLOCK_TIME_FORMAT, timeStr)
	if err != nil {
		return err
	}

	*this = BlockTime{Time: d}
	return nil
}

func (this BlockTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if this.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return this.Time, nil
}

func (this *BlockTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*this = BlockTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to BlockTime\n", v)
}

func (this BlockTime) ToString() string {
	return this.Format(BLOCK_TIME_FORMAT)
}

func (b *BlockTime) FromString(timeStr string) error {
	t, err := time.Parse(BLOCK_TIME_FORMAT, timeStr)
	if err != nil {
		return err
	}
	b.Time = t
	return nil
}
package mysql_server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type mysql struct {
	sql  *sql.DB
	gorm *gorm.DB
}

func InitializeGorm(database string) *mysql {
	dataSourceName := fmt.Sprintf("root:zJY121123!@tcp(127.0.0.1:3306)/"+
		"%s?charset=utf8mb4&parseTime=true&loc=Local&parseTime=true", database)
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
	gormDefaultDB.LogMode(true)
	fmt.Println("connect sql success")
	return &mysql{
		sql:  MySQLInlineActionClient,
		gorm: gormDefaultDB.LogMode(true),
	}
}

func Close(db *mysql) {
	CheckError(db.sql.Close, db.gorm.Close)
}

func CheckError(callBacks ...func() error) {
	for _, callBack := range callBacks {
		if err := callBack(); err != nil {
			fmt.Println(err)
		}
	}
}

type BlockTime struct {
	time.Time
}

type BatchInsertSqlBuilder = func(tableName, fields, valuePlaceholders string) string

func BatchInsertWithRawSql(db *gorm.DB, objArr []interface{}, builder BatchInsertSqlBuilder) error {
	fmt.Println(objArr)
	if len(objArr) == 0 {
		return nil
	}
	mainObj := objArr[0]
	mainScope := db.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))
	for i := range mainFields {
		// If primary key has blank value (0 for int, "" for string, nil for interface ...), skip it.
		// If field is ignore field, skip it.
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) || (mainFields[i].Relationship != nil) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
	}

	placeholdersArr := make([]string, 0, len(objArr))

	for _, obj := range objArr {
		scope := db.NewScope(obj)
		fields := scope.Fields()
		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) || (fields[i].Relationship != nil) {
				continue
			}
			placeholders = append(placeholders, scope.AddToVars(fields[i].Field.Interface()))
		}
		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
		// add real variables for the replacement of placeholders' '?' letter later.
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}

	rawSql := builder(
		mainScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "),
	)

	mainScope.Raw(rawSql)
	fmt.Println("sql cmd:", mainScope.SQLVars)
	result, err := mainScope.SQLDB().Exec(mainScope.SQL, mainScope.SQLVars...)
	if err != nil {
		fmt.Println("fail to exec batch insert. err:", err)
		fmt.Println("fail to exec batch insert. sql:", rawSql)
		return err
	}
	row, _ := result.RowsAffected()
	if row > 0 {
		//Logger.Debug("mainScope.TableName():%s,ROW:%d,quoted=%v,sql_vars:%v\n",mainScope.TableName(),row,quoted,mainScope.SQLVars)
	}
	return nil
}

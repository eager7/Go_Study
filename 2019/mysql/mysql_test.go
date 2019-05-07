package mysql

import (
	"fmt"
	"testing"
)

var table = `
CREATE TABLE IF NOT EXISTS eos_park.t_block_info
(
    block_num          INT UNSIGNED(11) NOT NULL COMMENT '区块编号',
    id                 CHAR(64) NOT NULL COMMENT '区块hash',
    producer           VARCHAR(12) NOT NULL COMMENT '出块节点',
    block_timestamp    DATETIME NOT NULL COMMENT '出块时间',
    transaction_mroot  CHAR(64) NOT NULL COMMENT 'transaction_mroot',
    ction_mroot        CHAR(64) NOT NULL COMMENT 'action_mroot',
    producer_signature VARCHAR(255) NOT NULL COMMENT '出块节点签名',
    transaction_num    INT UNSIGNED(8) NOT NULL DEFAULT '-1' COMMENT '交易数',
    action_num         INT UNSIGNED(8) NOT NULL DEFAULT '-1' COMMENT '操作数',
    PRIMARY KEY (block_num),
    UNIQUE KEY uk_id (id)
) ENGINE=INNODB;
`

func TestInitializeDB(t *testing.T) {
	db, err := InitializeDB()
	check(err)
	check(db.CreateTable(TableString))
	check(db.Insert())
	db.Show()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestMysqlFormat(t *testing.T) {
	fmt.Println(MysqlFormat(`Founder's" Token_0.1`))
}
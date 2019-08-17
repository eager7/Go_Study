package gorm

import "github.com/jinzhu/gorm"

type BidNameAction struct {
	GlobalSequence uint64 `gorm:"primary_key" json:"global_sequence"`
	Bidder         string `json:"bidder"`
	NewName        string `gorm:"column:newname" json:"newname"`
	Bid            string `json:"bid"`
	State          uint8  `json:"state"`
	ActionType     uint8  `json:"action_type"`
	TrxTimestamp   string `json:"trx_timestamp"`
	TransactionID  string `json:"transaction_id"`
	BlockNum       uint32 `json:"block_num"`
}

func (BidNameAction) TableName() string {
	return "t_bidname_action_info"
}

func Update(db *gorm.DB, b *BidNameAction) {
	db.Save(b)
}
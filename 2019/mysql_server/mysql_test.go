package mysql_server

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type ContractStatInfo struct {
	Index                int64 `gorm:"primary_key"`
	Account              string
	Date                 string
	InvokerNum           int
	InvokedTimes         int
	InvokerNum7Days      int     `gorm:"column:invoker_num_7d"`
	InvokedTimes7Days    int     `gorm:"column:invoked_times_7d"`
	EOSTransferAmount24h float64 `json:"eos_transfer_amount_24h" gorm:"column:eos_transfer_amount_24h"`
	TradeVolume24h       float64 `gorm:"column:trade_volume_24h"`
}

func (ContractStatInfo) TableName() string {
	return "t_contract_stat_info"
}

type contractStat struct {
	ContractName      string               // account 就是 contract name
	InvokedTimes      int                  // 被调用次数
	Invokers          map[string]time.Time // 调用者集合
	EosTransferAmount float64              // EOS 转入转出总额
	TradeVolume       float64              // 交易量
}

type tStatMap = map[string]contractStat

func TestContractStatInfoBatch(t *testing.T) {
	db := InitializeGorm("tTable")
	stats := make(map[string]contractStat)
	//allStats["test"] = map[string]contractStat{
	//	"statMap": {
	//		ContractName:      "contractStat",
	//		InvokedTimes:      0,
	//		Invokers:          map[string]time.Time{"pct": time.Now()},
	//		EosTransferAmount: 10,
	//		TradeVolume:       20,
	//	},
	//}
	jsonString := `{"betdiceadmin":{"ContractName":"betdiceadmin","InvokedTimes":5,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"betdicegroup":{"ContractName":"betdicegroup","InvokedTimes":2,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"betdicelotto":{"ContractName":"betdicelotto","InvokedTimes":2,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"betdicestake":{"ContractName":"betdicestake","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"betdicetasks":{"ContractName":"betdicetasks","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"betdividends":{"ContractName":"betdividends","InvokedTimes":4,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"betsandbacca":{"ContractName":"betsandbacca","InvokedTimes":1,"Invokers":{},"EosTransferAmount":1.2,"TradeVolume":0},"bgbetwallet1":{"ContractName":"bgbetwallet1","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"bostonshrimp":{"ContractName":"bostonshrimp","InvokedTimes":0,"Invokers":{},"EosTransferAmount":2,"TradeVolume":0},"chessiotoken":{"ContractName":"chessiotoken","InvokedTimes":3,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"dappbaccarat":{"ContractName":"dappbaccarat","InvokedTimes":0,"Invokers":{},"EosTransferAmount":0.4,"TradeVolume":0},"depostoken11":{"ContractName":"depostoken11","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"dragonoption":{"ContractName":"dragonoption","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"dtheoschain1":{"ContractName":"dtheoschain1","InvokedTimes":3,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"endlesstoken":{"ContractName":"endlesstoken","InvokedTimes":2,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosbetbank11":{"ContractName":"eosbetbank11","InvokedTimes":8,"Invokers":{},"EosTransferAmount":0.0002,"TradeVolume":0},"eosbetdice11":{"ContractName":"eosbetdice11","InvokedTimes":5,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosbettokens":{"ContractName":"eosbettokens","InvokedTimes":4,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosbiggame55":{"ContractName":"eosbiggame55","InvokedTimes":2,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosdotaprod1":{"ContractName":"eosdotaprod1","InvokedTimes":7,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eoshashdices":{"ContractName":"eoshashdices","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosio.null":{"ContractName":"eosio.null","InvokedTimes":3,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosjackscoin":{"ContractName":"eosjackscoin","InvokedTimes":6,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosjacksdice":{"ContractName":"eosjacksdice","InvokedTimes":2,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosjacksjack":{"ContractName":"eosjacksjack","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0.1038,"TradeVolume":0},"eosjackslead":{"ContractName":"eosjackslead","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosjoygame1b":{"ContractName":"eosjoygame1b","InvokedTimes":0,"Invokers":{},"EosTransferAmount":0.0956,"TradeVolume":0},"eosjoyiocoin":{"ContractName":"eosjoyiocoin","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosknightsio":{"ContractName":"eosknightsio","InvokedTimes":6,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosluckycoin":{"ContractName":"eosluckycoin","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosmax1token":{"ContractName":"eosmax1token","InvokedTimes":5,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosmaxiobull":{"ContractName":"eosmaxiobull","InvokedTimes":2,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosmaxiodraw":{"ContractName":"eosmaxiodraw","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosmaxioslot":{"ContractName":"eosmaxioslot","InvokedTimes":2,"Invokers":{},"EosTransferAmount":2.1364,"TradeVolume":0},"eosmaxioteam":{"ContractName":"eosmaxioteam","InvokedTimes":3,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosplaybrand":{"ContractName":"eosplaybrand","InvokedTimes":2,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eospokeriniu":{"ContractName":"eospokeriniu","InvokedTimes":2,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eostgctoken1":{"ContractName":"eostgctoken1","InvokedTimes":13,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"eosyxtoken11":{"ContractName":"eosyxtoken11","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"ezeosaccount":{"ContractName":"ezeosaccount","InvokedTimes":6,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"ffgametongbi":{"ContractName":"ffgametongbi","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"godice.e":{"ContractName":"godice.e","InvokedTimes":0,"Invokers":{},"EosTransferAmount":0.2,"TradeVolume":0},"higoldtokens":{"ContractName":"higoldtokens","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"houseaccount":{"ContractName":"houseaccount","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0.3629,"TradeVolume":0},"llgcontract1":{"ContractName":"llgcontract1","InvokedTimes":3,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"llgcontract2":{"ContractName":"llgcontract2","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"llgonebtotal":{"ContractName":"llgonebtotal","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"murmurdappco":{"ContractName":"murmurdappco","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"pornhashbaby":{"ContractName":"pornhashbaby","InvokedTimes":31,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"prochaintech":{"ContractName":"prochaintech","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"roulettespin":{"ContractName":"roulettespin","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"trustdicelog":{"ContractName":"trustdicelog","InvokedTimes":1,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"trustdicewin":{"ContractName":"trustdicewin","InvokedTimes":3,"Invokers":{},"EosTransferAmount":0.4238,"TradeVolume":0},"vsvscontract":{"ContractName":"vsvscontract","InvokedTimes":10,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"vsvsvsbetbet":{"ContractName":"vsvsvsbetbet","InvokedTimes":16,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0},"vsvsvsvipvip":{"ContractName":"vsvsvsvipvip","InvokedTimes":8,"Invokers":{},"EosTransferAmount":0,"TradeVolume":0}}`
	if err := json.Unmarshal([]byte(jsonString), &stats); err != nil {
		t.Fatal(err)
	}

	statInfoList := make([]interface{}, 0, 120)
		for _, statInfo := range stats {
			statInfoList = append(statInfoList, ContractStatInfo{
				Account:              statInfo.ContractName,
				Date:                 "",
				InvokedTimes:         statInfo.InvokedTimes,
				EOSTransferAmount24h: statInfo.EosTransferAmount,
				TradeVolume24h:       statInfo.TradeVolume,
			})
		}

	f := func(tableName, fields, valuePlaceholders string) string {
		cmd := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s ON DUPLICATE KEY UPDATE "+
			"`invoked_times`=`invoked_times`+VALUES(`invoked_times`), "+
			"`eos_transfer_amount_24h`=`eos_transfer_amount_24h`+VALUES(`eos_transfer_amount_24h`), "+
			"`trade_volume_24h`=`trade_volume_24h`+VALUES(`trade_volume_24h`)",
			tableName, fields, valuePlaceholders)
		return cmd
	}
	if err := BatchInsertWithRawSql(db.gorm, statInfoList, f); err != nil {
		t.Fatal(err)
	}
}

func TestContractStatInfoInsert(t *testing.T) {
	db := InitializeGorm("eos_park")
	c := ContractStatInfo{
		Index:                100,
		Account:              fmt.Sprintf("account[%d]", 100),
		Date:                 time.Now().Format("20060102"),
		InvokerNum:           0,
		InvokedTimes:         0,
		InvokerNum7Days:      0,
		InvokedTimes7Days:    0,
		EOSTransferAmount24h: 0,
		TradeVolume24h:       0,
	}
	if err := db.gorm.Save(c).Error; err != nil {
		t.Fatal(err)
	}
}

func TestSearch(t *testing.T) {
	// 查询过去七天的统计数据
	type tStat struct {
		InvokedTimes         int     `json:"invoked_times"`
		InvokerNum           int     `json:"invoker_num"`
		EOSTransferAmount24h float64 `json:"eos_transfer_amount_24h"`
		Date                 string  `json:"date"`
	}
	db := InitializeGorm("eos_park")
	stats := make([]tStat, 0)
	today := time.Now().Format("20060102")
	if err := db.gorm.
		Model(&ContractStatInfo{}).
		Select("`invoked_times`, `invoker_num`, `date`, `eos_transfer_amount_24h`").
		Where("`account` = ? AND `date` <= ?", "eosbluejacks", today).
		Order("`index` DESC").
		Limit(7).
		Scan(&stats).Error; err != nil {
		t.Fatal(err)
	}
	t.Log(stats)
}

package mysql_server

import (
	"fmt"
	"strconv"
	"strings"
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

func TestContractStatInfoBatch(t *testing.T) {
	db := InitializeGorm("eos_park")

	text := ` {0 eosbluejacks 20190325 0 1 0 0 0.01 0} {0 eoschessteam 20190325 0 1 0 0 0 0} {0 eosjackscoin 20190325 0 6 0 0 0 0} {0 eosbiggame22 20190325 0 1 0 0 0 0} {0 llgcontract2 20190325 0 1 0 0 0 0} {0 llgonebtotal 20190325 0 1 0 0 0 0} {0 endlesstoken 20190325 0 4 0 0 0 0} {0 eosdotaprod1 20190325 0 2 0 0 0 0} {0 eosbaccarat1 20190325 0 1 0 0 8 0} {0 bgbgbgbgbgbg 20190325 0 4 0 0 0 0}{0 eostgctoken1 20190325 0 13 0 0 0 0} {0 whaleextrust 20190325 0 3 0 0 0 0} {0 chessiotoken 20190325 0 4 0 0 0 0} {0 eosio.null 20190325 0 1 0 0 0 0} {0 endlessdicex 20190325 0 1 0 0 0.1 0} {0 eosjackslead 20190325 0 2 0 0 0 0} {0 leekdaogroup 20190325 0 1 0 0 0 0} {0betsacetoken 20190325 0 1 0 0 0 0} {0 tgoncomeon11 20190325 0 3 0 0 0 0} {0 betdicetasks 20190325 0 2 0 0 0 0} {0 eosknightsio 201903250 5 0 0 0 0} {0 eosmaxiobull 20190325 0 2 0 0 0 0} {0 trustdicelog 20190325 0 1 0 0 0 0} {0 eosbiggame55 20190325 0 2 0 0 10.161000000000001 0} {0 eosbetdice11 20190325 0 5 0 0 0 0} {0 baccarat.e 20190325 0 2 0 0 0 0} {0 vsvscontract 20190325 0 8 0 0 0 0} {0 eospokeriniu 20190325 0 1 0 0 0 0} {0 betdicelucky 20190325 0 1 0 0 0 0} {0 pokereosbull 20190325 0 0 0 0 8 0} {0 vsvsvsbetbet 20190325 0 7 0 0 0.1289 0} {0 godice.e 20190325 0 4 0 0 0 0} {0 eostowergame 20190325 0 1 0 0 0 0} {0 eosnowbetext 20190325 0 0 0 0 0.01 0} {0 eoschessdice 20190325 0 2 0 0 1 0} {0 bgbetwallet1 20190325 0 1 0 0 0 0} {0 eoshashdices 20190325 0 1 0 0 0 0} {0 trustdicewin 20190325 0 1 0 0 1.888 0} {0 betdicetoken 20190325 0 6 0 0 0 0} {0 eosbetbank11 20190325 0 5 0 0 0.40630000000000005 0} {0 pornhashbaby 20190325 0 24 0 0 00} {0 huobideposit 20190325 0 0 0 0 1053 0} {0 lynxtoken123 20190325 0 1 0 0 0 0} {0 ffgameniuniu 20190325 0 1 0 0 0 0} {0 vsvsvsvipvip 20190325 0 4 0 0 0 0} {0 biggamerefer 20190325 0 1 0 0 0 0} {0 newdexpocket 20190325 0 0 0 0 0.0118 0} {0 bntbntbntbnt 20190325 0 2 0 0 00} {0 eosjacksjack 20190325 0 6 0 0 0 0} {0 everipediaiq 20190325 0 5 0 0 0 0} {0 fuckeoscpuuu 20190325 0 1 0 0 0 0} {0 findexfindex 20190325 0 2 0 0 4.4292 0} {0 ezeosaccount 20190325 0 3 0 0 1 0} {0 betdicegroup 20190325 0 4 0 0 0 0} {0 biggameminer 20190325 0 1 0 0 0 0} {0 thisisbancor 20190325 0 0 0 0 40.4 0} {0 betdiceadmin 20190325 0 7 0 0 0 0} {0 endlesslogs1 20190325 0 2 0 0 0 0} {0 houseaccount 20190325 0 3 0 0 0 0} {0 betdicelotto 20190325 0 2 0 0 0 0} {0 betdividends 20190325 0 2 0 0 0 0}`
	text = strings.Replace(text, " {", "", -1)
	list := strings.Split(text, "}")
	statInfoList := make([]interface{}, 0)
	for _, l := range list {
		objs := strings.Split(l, " ")
		fmt.Println("objs:", len(objs), objs)
		if len(objs) != 9 {
			continue
		}
		//invokerNum, err := strconv.ParseInt(objs[3], 10, 32)
		//if err != nil {
		//	t.Fatal(err)
		//}
		invokerTimes, err := strconv.ParseInt(objs[4], 10, 32)
		if err != nil {
			t.Fatal(err)
		}
		//invokerNum7d, err := strconv.ParseInt(objs[5], 10, 32)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//invokerTimes7d, err := strconv.ParseInt(objs[6], 10, 32)
		//if err != nil {
		//	t.Fatal(err)
		//}
		eosTransferAmount, err := strconv.ParseFloat(objs[7], 64)
		if err != nil {
			t.Fatal(err)
		}
		tradeVolume24, err := strconv.ParseFloat(objs[8], 64)
		if err != nil {
			t.Fatal(err)
		}
		statInfoList = append(statInfoList, ContractStatInfo{
			Account:              objs[1],
			Date:                 time.Now().Format("20060102"),
			InvokedTimes:         int(invokerTimes),
			EOSTransferAmount24h: float64(eosTransferAmount),
			TradeVolume24h:       float64(tradeVolume24),
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
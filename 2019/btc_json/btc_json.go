package js

import (
	"errors"
	"strconv"
)

const (
	Decimals      = 100000000
	initReward    = 50
	firstHalving  = 210000 //此区块后奖励为25BTC
	secondHalving = 420000 //此区块后奖励为12.5BTC
	thirdHalving  = 630000 //此区块后奖励为6.25BTC
)

var ErrCatchMain = errors.New("catch main")

func Reward(height uint64) float64 {
	if height < firstHalving {
		return initReward
	}
	if height < secondHalving {
		return initReward / 2
	}
	if height < thirdHalving {
		return initReward / 4
	}
	return initReward / 8
}

func StringToFloat64(v string) (float64, error) {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, errors.New("StringToFloat64 err:" + err.Error())
	}
	return f, nil
}
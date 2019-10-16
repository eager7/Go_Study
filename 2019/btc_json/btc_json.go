package js

import "errors"

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

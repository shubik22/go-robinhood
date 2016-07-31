package app

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/shubik22/robinhood-client"
)

type RobinhoodCache struct {
	cache *cache.Cache
}

func NewCache() *RobinhoodCache {
	cache := cache.New(cache.NoExpiration, cache.NoExpiration)
	return &RobinhoodCache{cache: cache}
}

func (rc *RobinhoodCache) SetAccount(username string, account *robinhood.Account) {
	key := fmt.Sprintf("%v:account", username)
	rc.cache.Set(key, *account, cache.NoExpiration)
}

func (rc *RobinhoodCache) GetAccount(username string) robinhood.Account {
	key := fmt.Sprintf("%v:account", username)
	a, _ := rc.cache.Get(key)
	return a.(robinhood.Account)
}

func (rc *RobinhoodCache) SetPositions(username string, pr *robinhood.PositionsResponse) {
	key := fmt.Sprintf("%v:positions", username)
	rc.cache.Set(key, *pr, cache.NoExpiration)
}

func (rc *RobinhoodCache) GetPositions(username string) robinhood.PositionsResponse {
	key := fmt.Sprintf("%v:positions", username)
	p, _ := rc.cache.Get(key)
	return p.(robinhood.PositionsResponse)
}

func (rc *RobinhoodCache) SetQuote(positionUrl string, quote *robinhood.Quote) {
	key := fmt.Sprintf("%v:quote", positionUrl)
	rc.cache.Set(key, *quote, cache.NoExpiration)
}

func (rc *RobinhoodCache) GetQuote(positionUrl string) robinhood.Quote {
	key := fmt.Sprintf("%v:quote", positionUrl)
	a, _ := rc.cache.Get(key)
	return a.(robinhood.Quote)
}

func (rc *RobinhoodCache) SetLeaderboard(lb *robinhood.Leaderboard) {
	rc.cache.Set("leaderboard", *lb, cache.NoExpiration)
}

func (rc *RobinhoodCache) GetLeaderboard() robinhood.Leaderboard {
	lb, _ := rc.cache.Get("leaderboard")
	return lb.(robinhood.Leaderboard)
}

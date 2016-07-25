package app

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/shubik22/go-robinhood/lib/models"
)

type RobinhoodCache struct {
	cache *cache.Cache
}

func NewCache() *RobinhoodCache {
	cache := cache.New(cache.NoExpiration, cache.NoExpiration)
	return &RobinhoodCache{cache: cache}
}

func (rc *RobinhoodCache) SetAccount(username string, account *models.Account) {
	key := fmt.Sprintf("%v:account", username)
	rc.cache.Set(key, *account, cache.NoExpiration)
}

func (rc *RobinhoodCache) GetAccount(username string) models.Account {
	key := fmt.Sprintf("%v:account", username)
	a, _ := rc.cache.Get(key)
	return a.(models.Account)
}

func (rc *RobinhoodCache) SetPositions(username string, pr *models.PositionsResponse) {
	key := fmt.Sprintf("%v:positions", username)
	rc.cache.Set(key, *pr, cache.NoExpiration)
}

func (rc *RobinhoodCache) GetPositions(username string) models.PositionsResponse {
	key := fmt.Sprintf("%v:positions", username)
	p, _ := rc.cache.Get(key)
	return p.(models.PositionsResponse)
}

func (rc *RobinhoodCache) SetQuote(positionUrl string, quote *models.Quote) {
	key := fmt.Sprintf("%v:quote", positionUrl)
	rc.cache.Set(key, *quote, cache.NoExpiration)
}

func (rc *RobinhoodCache) GetQuote(positionUrl string) models.Quote {
	key := fmt.Sprintf("%v:quote", positionUrl)
	a, _ := rc.cache.Get(key)
	return a.(models.Quote)
}

func (rc *RobinhoodCache) SetLeaderboard(lb *models.Leaderboard) {
	rc.cache.Set("leaderboard", *lb, cache.NoExpiration)
}

func (rc *RobinhoodCache) GetLeaderboard() models.Leaderboard {
	lb, _ := rc.cache.Get("leaderboard")
	return lb.(models.Leaderboard)
}

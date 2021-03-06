package app

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/shubik22/robinhood-client"
	"github.com/shubik22/rbhd/lib/managers"
	"math/rand"
	"sort"
	"strconv"
)

type App struct {
	cache *RobinhoodCache
	cm    *managers.ClientManager
}

func NewApp() *App {
	a := &App{
		cache: NewCache(),
		cm:    managers.NewClientManager(),
	}
	return a
}

func (a *App) Init() {
	a.populateCache()
	a.calculateLeaderboard()
	a.setupCronJobs()
}

func (a *App) populateCache() {
	for _, c := range a.cm.Clients {
		fmt.Printf("Fetching data for %v\n", c.UserName)
		a.FetchData(c)
	}
}

func (a *App) setupCronJobs() {
	cr := cron.New()

	for _, c := range a.cm.Clients {
		minutes := rand.Intn(10) + 20
		seconds := rand.Intn(60)
		durationStr := fmt.Sprintf("@every %vm%vs", minutes, seconds)
		fmt.Printf("Setting cron job to update %v %v\n", c.UserName, durationStr)
		cr.AddFunc(durationStr, func() {
			a.FetchData(c)
			a.calculateLeaderboard()
		})
	}

	cr.Start()
}

func (a *App) FetchData(c *robinhood.Client) {
	a.FetchAccounts(c)
	a.FetchPositionsAndQuotes(c)
}

func (a *App) FetchAccounts(c *robinhood.Client) {
	fmt.Printf("Fetching account for %v\n", c.UserName)
	ar, _, err := c.Accounts.ListAccounts()
	if err != nil {
		fmt.Printf("Failed fetching accounts for %v.\n", c.UserName)
		return
	} else if len(ar.Results) == 0 {
		fmt.Printf("Found zero accounts for %v.\n", c.UserName)
		return
	} else if len(ar.Results) > 1 {
		fmt.Printf("Found %v accounts for %v.\n", len(ar.Results), c.UserName)
		return
	}

	a.cache.SetAccount(c.UserName, &ar.Results[0])
}

func (a *App) FetchPositionsAndQuotes(c *robinhood.Client) {
	fmt.Printf("Fetching positions for %v\n", c.UserName)
	pr, _, err := c.Positions.ListPositions()
	if err != nil {
		fmt.Printf("Failed fetching positions for %v\nError: %v", c.UserName, err)
		return
	}

	for _, p := range pr.Results {
		q, _ := strconv.ParseFloat(p.Quantity, 64)
		if q == 0 {
			continue
		}

		quote, _, err := c.Quotes.GetQuote(&p)
		if err != nil {
			fmt.Printf("Failed fetching quote for %v\nError: %v", p.Instrument, err)
			continue
		}

		a.cache.SetQuote(p.URL, quote)
	}

	a.cache.SetPositions(c.UserName, pr)
}

func (a *App) GetLeaderboard() *robinhood.Leaderboard {
	lb := a.cache.GetLeaderboard()
	return &lb
}

func (a *App) calculateLeaderboard() {
	lb := &robinhood.Leaderboard{}
	var users []robinhood.User
	for _, c := range a.cm.Clients {
		user := a.buildUser(c.UserName)
		if user != nil {
			users = append(users, *user)
		}
	}

	sort.Sort(ByTotalBalance(users))
	lb.Entries = users
	a.cache.SetLeaderboard(lb)
}

func (a *App) buildUser(username string) *robinhood.User {
	u := &robinhood.User{
		Username: username,
		Positions: make([]robinhood.SimplePosition, 0),
	}
	account := a.cache.GetAccount(username)
	cashBalance, _ := strconv.ParseFloat(account.BuyingPower, 64)
	u.CashBalance = cashBalance

	var positionBalance float64
	pr := a.cache.GetPositions(username)
	for _, p := range pr.Results {
		quantity, _ := strconv.ParseFloat(p.Quantity, 64)
		if quantity == 0 {
			continue
		}

		q := a.cache.GetQuote(p.URL)
		lastPrice, _ := strconv.ParseFloat(q.LastTradePrice, 64)
		avgBuyPrice, _ := strconv.ParseFloat(p.AverageBuyPrice, 64)
		simplePosition := robinhood.SimplePosition{
			PurchaseTime:    p.CreatedAt,
			Quantity:        quantity,
			Symbol:          q.Symbol,
			AverageBuyPrice: avgBuyPrice,
			LastTradePrice:  lastPrice,
		}
		u.Positions = append(u.Positions, simplePosition)
		positionBalance += (quantity * lastPrice)
	}

	totalBalance := cashBalance + positionBalance
	u.PositionBalance = positionBalance
	u.TotalBalance = totalBalance

	return u
}

type ByTotalBalance []robinhood.User

func (a ByTotalBalance) Len() int {
	return len(a)
}

func (a ByTotalBalance) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByTotalBalance) Less(i, j int) bool {
	return a[i].TotalBalance < a[j].TotalBalance
}

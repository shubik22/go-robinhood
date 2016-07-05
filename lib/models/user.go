package models

type User struct {
	Username        string  `json:"username"`
	CashBalance     float64 `json:"cash_balance"`
	PositionBalance float64 `json:"position_balance"`
}

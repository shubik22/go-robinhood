package managers

import (
	"github.com/shubik22/robinhood"
)

type ClientManager struct {
	Clients []*robinhood.Client
}

func NewClientManager() *ClientManager {
	credentials := NewCredentialsManager()
	cm := &ClientManager{}
	for _, u := range credentials.AllUsers() {
		pw, err := credentials.GetPassword(u)
		if err == nil {
			c := NewClient(u, pw)
			cm.Clients = append(cm.Clients, c)
		}
	}

	return cm
}

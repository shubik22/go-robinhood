package managers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	credentialsFile = "config/credentials.json"
)

type userStore struct {
	Users []userCredentials `json:"users"`
}

type userCredentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type CredentialsManager struct {
	UsersMap map[string]string
}

func NewCredentialsManager() *CredentialsManager {
	b, err := ioutil.ReadFile(credentialsFile)

	if err != nil {
		panic(err)
	}

	var us userStore

	if err := json.Unmarshal(b, &us); err != nil {
		panic(err)
	}

	cm := &CredentialsManager{
		UsersMap: make(map[string]string),
	}

	for _, el := range us.Users {
		cm.AddUser(el)
	}

	return cm
}

func (c *CredentialsManager) AddUser(u userCredentials) {
	c.UsersMap[u.UserName] = u.Password
}

func (c *CredentialsManager) AllUsers() []string {
	users := make([]string, 0, len(c.UsersMap))
	for u := range c.UsersMap {
		users = append(users, u)
	}
	return users
}

func (c *CredentialsManager) GetPassword(u string) (string, error) {
	pw := c.UsersMap[u]
	if pw == "" {
		return "", fmt.Errorf("Missing password for %v", u)
	}

	return pw, nil
}

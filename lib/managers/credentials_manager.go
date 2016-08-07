package managers

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type CredentialsManager struct {
	UsersMap map[string]string
}

func NewCredentialsManager() *CredentialsManager {
	godotenv.Load()

	cm := &CredentialsManager{
		UsersMap: make(map[string]string),
	}

	usersStr := os.Getenv("USERNAMES")
	usernames := strings.Split(usersStr, ",")

	for _, username := range usernames {
		pwKey := fmt.Sprintf("%v:pw", username)
		pw := os.Getenv(pwKey)
		cm.UsersMap[username] = pw
	}

	return cm
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

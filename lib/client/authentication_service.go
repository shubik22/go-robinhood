package client

import (
	"net/http"
	"net/url"
)

type authResponse struct {
	Token string `json:"token"`
}

type authRequest struct {
	username string `json:"username"`
	password string `json:"password"`
}

type AuthenticationService service

func (s *AuthenticationService) Login() (*http.Response, error) {
	username := s.client.UserName
	password, err := s.client.Credentials.GetPassword(username)

	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("username", username)
	params.Add("password", password)

	a := &authResponse{}
	resp, err := s.client.PostForm("api-token-auth/", params, a)

	if err == nil {
		s.client.AuthToken = a.Token
	}

	return resp, err
}

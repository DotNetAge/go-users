package services

import "fmt"

type AccessToken struct {
	Value        string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int16  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	UID          string `json:"uid"`
	IID          string `json:"iid"`
	Roles        string `json:"roles"`
}

func (a *AccessToken) ToString() string {
	tmpl := "?access_token=%s&token_type=%s&expires_in=%6d&refresh_token=%s&uid=%s&iid=%s&roles=%s"
	return fmt.Sprintf(tmpl, a.Value, a.TokenType, a.ExpiresIn, a.RefreshToken, a.UID, a.IID, a.Roles)
}

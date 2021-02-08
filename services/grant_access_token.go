package services

import (
	//"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"go-uds/models"
	"time"
)

type accessClaims struct {
	UserName string `json:"name"`
	UserID   string `json:"uid"`
	IID      string `json:"iid"`
	Roles    string `json:"roles,omitempty"`
	Scope    string `json:"scope,omitempty"`
}

func (s *authService) GrantAccessToken(identityID, userID, userName, scope string, roles string) (*AccessToken, error) {
	// 使用client_secret的原因是因为考虑到应用上下文传递的问题，如果将secret由的用户目录产生那将必然会导致每个
	// 微服务中都需要该secret对JWT进行解码，这就会形成一个很大的依赖性存在。
	claims := accessClaims{
		userName,
		userID,
		identityID,
		roles,
		scope,
	}

	refreshClaims := jwt.Claims{Subject: identityID}

	tokenPair, err := s.Context.Signer.NewTokenPair(claims, refreshClaims, 23*24*time.Hour)

	if err != nil {
		return nil, err
	}

	err = s.Context.RefreshTokens.Save(&models.RefreshToken{
		string(tokenPair.RefreshToken),
		userName,
		userID,
		identityID,
		roles,
		scope,
	})

	if err != nil {
		return nil, err
	}

	return &AccessToken{
		Value:        tokenPair.AccessToken,
		TokenType:    "Bear",
		ExpiresIn:    3600,
		RefreshToken: tokenPair.RefreshToken,
		UID:          userID,
		IID:          identityID,
		Roles:        roles,
	}, nil
}

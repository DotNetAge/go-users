package services

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/iris-contrib/go.uuid"
)

type UserClaims struct {
	UserName string `json:"name"`
	UserID   string `json:"uid"`
	IID      string `json:"iid"`
	Roles    string `json:"roles"`
	Scope    string `json:"scope"`
	jwt.StandardClaims
}

func (s *authService) GrantAccessToken(client_id, client_secret, identity_id, user_id, user_name, scope string, roles string) (*AccessToken, error) {
	// 使用client_secret的原因是因为考虑到应用上下文传递的问题，如果将secret由的用户目录产生那将必然会导致每个
	// 微服务中都需要该secret对JWT进行解码，这就会形成一个很大的依赖性存在。
	claims := UserClaims{
		user_name,
		user_id,
		identity_id,
		roles,
		scope,
		jwt.StandardClaims{
			ExpiresAt: 3600, // exp
		},
	}
	// claims := make(map[string]interface{})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, err := token.SignedString(client_secret)
	if err != nil {
		return nil, err
	}
	refresh_token := uuid.Must(uuid.NewV1()).String()
	key := "refresh-token:" + refresh_token
	//TODO: 此处需要重构将仅需要的value写入Redis
	_, err = s.Context.Redis.HMSet(ctx, key,
		"name", user_name,
		"iid", identity_id,
		"uid", user_id,
		"scope", scope,
		"roles", roles).Result()

	if err != nil {
		return nil, err
	}
	return &AccessToken{
		Value:        access_token,
		TokenType:    "Bear",
		ExpiresIn:    3600,
		RefreshToken: refresh_token,
		UID:          user_id,
		IID:          identity_id,
		Roles:        roles,
	}, nil
}

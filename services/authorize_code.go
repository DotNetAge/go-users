package services

import (
	"context"
	"math/rand"
	"time"
)

var (
	ctx      = context.Background()
	runeMask = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStrNumber(n int, mask []rune) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = mask[rand.Intn(len(mask))]
	}
	return string(b)
}

// 生成临时授权码
func (s *authService) AuthorizeCode(clientId, mobile, scope string) (string, error) {
	code := RandStrNumber(4, runeMask)
	key := "authorize-codes:" + clientId + ":" + code
	timeout := 30 * time.Second
	_, err := s.Context.Redis.Set(ctx, key+":mobile", mobile, timeout).Result()

	if err != nil {
		return "", err
	}

	_, err = s.Context.Redis.Set(ctx, key+":scope", scope, timeout).Result()

	if err != nil {
		return "", err
	}

	return code, err
}

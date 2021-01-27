package services

import "github.com/go-redis/redis/v8"

// 生成的手机验证码
func (s *authService) GrantVerifyCode(clientId, mobile string) (*VerifyResponse, error) {
	// 检查该手机是否已经生成过验证码
	key := "verify-codes:" + clientId + ":" + mobile
	verify_code, err := s.Context.Redis.Get(ctx, key).Result()

	// 如果验证码不存在就生成一个并存入redis
	if err == redis.Nil {
		verify_code = RandStrNumber(4, []rune("1234567890"))
		err = s.Context.Redis.Set(ctx, key, verify_code, 20).Err()
		if err != nil {
			return nil, err
		}
	}

	return &VerifyResponse{
		Code: verify_code,
	}, nil
}

func (s *authService) ValidateVerifyCode(clientId, mobile, code string) bool {

	// 检查该手机是否已经生成过验证码
	key := "verify-codes:" + clientId + ":" + mobile
	verify_code, err := s.Context.Redis.Get(ctx, key).Result()

	// 如果验证码不存在就生成一个并存入redis
	if err == redis.Nil {
		return false
	}

	if verify_code == code {
		return true
	}

	return false

}

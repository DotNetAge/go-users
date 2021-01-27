package services

import (
	"github.com/go-redis/redis/v8"
	"reflect"
)

func (s *authService) GrantByRefreshToken(clientId, refreshToken string) (*AccessToken, error) {
	key := "refresh-token:" + refreshToken
	//claims := &UserClaims{}
	//t := reflect.TypeOf(claims)
	//v := reflect.ValueOf(claims)
	//
	//// 从 Redis 中将HMSet的数据读入至Claims中
	//for i := 0; i < t.NumField(); i++ {
	//	fieldName := t.Field(i).Name
	//	val, err := s.Context.Redis.HGet(ctx, key, fieldName).Result()
	//	if err != redis.Nil && err != nil {
	//		//TODO: 如果是exp字段则会导致出错，因为exp字段是整数
	//		if val != "" {
	//			v.Field(i).SetString(val)
	//		}
	//	}
	//}
	vals, err := s.Context.Redis.HMGet(ctx, key, "name","iid","uid","roles","scope").Result()

	// s.GrantAccessToken()
	// 这里使用HMGet 值得一讲其中的技巧
	return nil, nil
}

/* 反射示例
func main() {
    a := map[string]int{
        "A": 1, "B": 2,
    }
    keys := reflect.ValueOf(a).MapKeys()
    strkeys := make([]string, len(keys))
    for i := 0; i < len(keys); i++ {
        strkeys[i] = keys[i].String()
    }
    fmt.Print(strings.Join(strkeys, ","))
}
*/

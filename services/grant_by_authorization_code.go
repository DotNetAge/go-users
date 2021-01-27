package services

import "go-uds/models"

func (s *authService) GrantByAuthorizationCode(clientId, code string) (*AccessToken, error) {
	key := "authorize-codes:" + clientId + ":" + code
	fields, err := s.Context.Redis.HMGet(ctx, key, "mobile", "scope").Result()
	mobile := fields[0].(string)
	scope := fields[1].(string)
	if err != nil {
		return nil, err
	}
	clientSecret := ""
	identity, e := s.Context.Identities.FindByName(mobile)
	if e != nil {
		return nil, e
	}
	roles := ""

	if nil == identity {
		// create user
		user := models.NewUser(mobile)
		user.Mobile = mobile
		uid, _err := s.Context.Users.Save(user)
		if _err != nil {
			return nil, nil
		}
		identity = models.NewIdentity(mobile)
		identity.UID = uid
		_, _err = s.Context.Identities.Save(identity)
		if _err != nil {
			return nil, nil
		}
	}

	token, _err := s.GrantAccessToken(clientId, clientSecret, identity.ID, identity.UID, identity.Name, scope, roles)
	if _err != nil {
		return nil, nil
	}
	return token, nil
}

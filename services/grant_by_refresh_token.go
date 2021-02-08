package services

func (s *authService) GrantByRefreshToken(clientId, refreshToken string) (*AccessToken, error) {
	token, _err := s.Context.RefreshTokens.Find(refreshToken)

	if _err != nil {
		return nil, _err
	}

	_, _err = s.Context.RefreshTokens.Delete(refreshToken)

	return s.GrantAccessToken(token.IID, token.UserID, token.UserName, token.Scope, token.Roles)

}
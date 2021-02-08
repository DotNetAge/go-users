package services

import (
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12/middleware/jwt"
	"go-uds/repositories"
)

type VerifyResponse struct {
	Code string `json:"verify_code"`
}

type AuthService interface {
	AuthUser(username, password string) bool
	AuthorizeCode(clientID, mobile, scope string) (string, error)
	GrantVerifyCode(clientID, mobile string) (*VerifyResponse, error)
	GrantByAuthorizationCode(clientID, code string) (*AccessToken, error)
	GrantByRefreshToken(clientID, refreshToken string) (*AccessToken, error)
	GrantByPassword(clientID, username, password, scope string) (*AccessToken, error)
	GrantByClaims(clientID, name, source string, claims map[string]string) (*AccessToken, error)
	GrantAccessToken(identityID, userID, userName, scope string, roles string) (*AccessToken, error)
	ValidateVerifyCode(clientID, mobile, code string) bool
	ValidateClient(clientID string) bool
	ValidateMobile(mobile string) bool
}

type ServiceContext struct {
	Identities repositories.IdentityRepository
	Users      repositories.UserRepository
	RefreshTokens repositories.RefreshTokenRepository
	Signer     jwt.Signer
	Verifier   jwt.Verifier
	Redis      *redis.Client
}

type authService struct {
	Context *ServiceContext
}

func NewAuthService(c *ServiceContext) AuthService {
	return &authService{Context: c}
}

func (s *authService) ValidateClient(clientId string) bool {
	panic("Not implatement!")
	return false
}

func (s *authService) ValidateMobile(mobile string) bool {
	identity, _ := s.Context.Identities.FindByName(mobile)
	return identity != nil
}

func (s *authService) AuthUser(username, password string) bool {
	v, _ := s.Context.Identities.Valid(username, password)
	return v
}

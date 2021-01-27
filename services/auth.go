package services

import (
	"github.com/go-redis/redis/v8"
	"go-uds/repositories"
)

type VerifyResponse struct {
	Code string `json:"verify_code"`
}

type AuthService interface {
	AuthUser(username, password string) bool
	AuthorizeCode(clientId, mobile, scope string) (string, error)
	GrantVerifyCode(clientId, mobile string) (*VerifyResponse, error)
	GrantByAuthorizationCode(client_id, code string) (*AccessToken, error)
	GrantByRefreshToken(clientId, refreshToken string) (*AccessToken, error)
	GrantByPassword(client_id, username, password, scope string) (*AccessToken, error)
	GrantByClaims(client_id, name, source string, claims map[string]string) (*AccessToken, error)
	GrantAccessToken(client_id, client_secret, identity_id, user_id, user_name, scope string, roles string) (*AccessToken, error)
	ValidateVerifyCode(clientId, mobile, code string) bool
	ValidateClient(clientId string) bool
	ValidateMobile(mobile string) bool
}

type ServiceContext struct {
	Identities repositories.IdentityRepository
	Users      repositories.UserRepository
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

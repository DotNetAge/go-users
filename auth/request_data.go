package auth

type ResponseTypes string
type GrantTypes string

const (
	ResponseType_Code           ResponseTypes = "code"
	ResponseType_Token          ResponseTypes = "token"
	GrantType_AuthorizationCode GrantTypes    = "authorization_code"
	GrantType_Password          GrantTypes    = "password"
	GrantType_RefreshToken      GrantTypes    = "refresh_token"
	GrantType_Credentials       GrantTypes    = "credentials"
)

type VerifyData struct {
	ClientId string `json:"client_id"`
	Mobile   string `json:"mobile"`
	Code     string `json:"verify_code"`
}

type AuthData struct {
	ClientId     string        `form:"client_id" json:"client_id"`
	Mobile       string        `form:"mobile"`
	VerifyCode   string        `form:"verify_code"`
	ResponseType ResponseTypes `form:"response_type" json:"response_type"`
	RedirectURI  string        `form:"redirect_uri" json:"redirect_uri"`
	Scope        string        `form:"scope" json:"scope"`
	State        string        `form:"state" json:"state"`
}

type RefreshTokenData struct {
	ClientId     string `json:"client_id"`
	RefreshToken string `json:"refresh_token"`
}

type TokenData struct {
	ClientId  string     `json:"client_id"`
	GrantType GrantTypes `json:"grant_type"`
	Code      string     `json:"code"`
}

type ExternalCodeTokenData struct {
	ClientId  string     `json:"client_id"`
	GrantType GrantTypes `json:"grant_type"`
	Code      string     `json:"code"`
	AppId     string     `json:"appid"`
}

type PasswordTokenData struct {
	GrantType GrantTypes `json:"grant_type"`
	ClientId  string     `json:"client_id"`
	UserName  string     `json:"username"`
	Password  string     `json:"password"`
	Scope     string     `json:"scope"`
}

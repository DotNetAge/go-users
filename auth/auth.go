package auth

import (
	"github.com/kataras/iris/v12"
	"go-uds/services"
)

func RegisterRoutes(router *iris.APIContainer) {
	router.Get("/authorize", AuthorizeHandler)
	router.Post("/authorize", LoginHandler)
	router.Post("/verify", VerifyHandler)
	router.Post("/token", TokenHandler)
	router.Post("/token/{channel}", TokenExternalHandler)
}

func AuthorizeHandler(ctx iris.Context) {
	var authData AuthData
	// 从QueryString中读取数据
	if err := ctx.ReadQuery(&authData); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.ViewData("Title", "登录")
	ctx.View("login.html", authData)
}

func LoginHandler(ctx iris.Context, s services.AuthService) {
	var authData AuthData
	// 从Form中读取数据
	if err := ctx.ReadForm(&authData); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if s.ValidateVerifyCode(authData.ClientId, authData.Mobile, authData.VerifyCode) {

		if authData.ResponseType == ResponseType_Token {
			accessToken, err := s.GrantByMobile(authData.ClientId, authData.Mobile, authData.Scope)
			if err != nil {
				ctx.StopWithError(iris.StatusInternalServerError, err)
				return
			}
			redirectURI := authData.RedirectURI + "?" + accessToken.ToString()
			ctx.Redirect(redirectURI)

		} else {
			code, err := s.AuthorizeCode(authData.ClientId, authData.Mobile, authData.Scope)

			if err != nil {
				ctx.StopWithError(iris.StatusBadRequest, err)
				return
			}

			redirectURI := authData.RedirectURI + "?code=" + code + "&state=" + authData.State
			ctx.Redirect(redirectURI)
		}
	}

	//ctx.StopWithError(iris.StatusUnauthorized,error.Error)
}

func VerifyHandler(ctx iris.Context, s services.AuthService) {
	var verifyCode VerifyData
	if err := ctx.ReadJSON(&verifyCode); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	resp, err := s.GrantVerifyCode(verifyCode.ClientId, verifyCode.Mobile)
	if err != nil {
		ctx.StopWithError(iris.StatusInvalidToken, err)
		return
	}

	ctx.JSON(resp)
}

func TokenHandler(ctx iris.Context, s services.AuthService) {
	var accessToken interface{}
	var token TokenData
	err := ctx.ReadJSON(&token)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	switch token.GrantType {
	case GrantType_Credentials:
		ctx.StopWithStatus(iris.StatusNotImplemented)
		return
	case GrantType_Password:
		var passToken PasswordTokenData
		if err = ctx.ReadJSON(&passToken); err != nil {
			ctx.StopWithError(iris.StatusBadRequest, err)
			return
		}
		accessToken, err = s.GrantByPassword(passToken.ClientId, passToken.UserName, passToken.Password, passToken.Scope)
	case GrantType_RefreshToken:
		var refreshToken RefreshTokenData
		if err = ctx.ReadJSON(&refreshToken); err != nil {
			ctx.StopWithError(iris.StatusBadRequest, err)
			return
		}
		accessToken, err = s.GrantByRefreshToken(refreshToken.ClientId, refreshToken.RefreshToken)
	default:
		accessToken, err = s.GrantByAuthorizationCode(token.ClientId, token.Code)
	}

	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
	}

	if _, err = ctx.JSON(accessToken); err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
	}
}

func TokenExternalHandler(ctx iris.Context, s services.AuthService) {
	var extToken ExternalCodeTokenData
	if err := ctx.ReadJSON(&extToken); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	channel_name := ctx.Params().Get("channel")
	//TODO: get user claims and identity name
	var claims map[string]string
	name := "unnamed"

	token, err := s.GrantByClaims(extToken.ClientId, name, channel_name, claims)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	if _, err := ctx.JSON(token); err != nil {
		ctx.StopWithError(iris.StatusInvalidToken, err)
	}
}

package apihandlers

import (
	ctxvalue "medods-test/internal/ctx-value"
	dbhandlers "medods-test/internal/db-handlers"
	"medods-test/internal/tokens"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTokens(c echo.Context) error {
	cc := c.(*CustomContext)
	log := ctxvalue.GetLog(cc.Ctx)

	var tokenRequest TokenRequest
	if err := c.Bind(&tokenRequest); err != nil {
		log.Error("bind TokenRequest", "error", err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	if !dbhandlers.CheckUser(cc.Ctx, tokenRequest.GUID) {
		log.Debug("no such user")
		return c.String(http.StatusBadRequest, "no such user")
	}

	t, err := tokens.CreateAccessToken(tokenRequest.IPAddress, "secret")
	if err != nil {
		log.Error("create access token", "error", err)
		return c.String(http.StatusInternalServerError, "create token")
	}

	refreshToken, refreshTokenHashed, err := tokens.CreateRefreshToken()
	if err != nil {
		log.Error("create refresh token", "error", err)
		return c.String(http.StatusInternalServerError, "create refresh token")
	}

	log.Debug("login", "refresh token", refreshToken, "hash", string(refreshTokenHashed))

	if err := dbhandlers.UpdateRefreshToken(cc.Ctx, tokenRequest.GUID, tokenRequest.IPAddress, refreshTokenHashed); err != nil {
		log.Error("update refresh token", "error", err)
		return c.String(http.StatusInternalServerError, "update refresh token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access token":  t,
		"refresh token": refreshToken,
	})
}

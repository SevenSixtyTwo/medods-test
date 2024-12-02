package apihandlers

import (
	ctxvalue "medods-test/internal/ctx-value"
	dbhandlers "medods-test/internal/db-handlers"
	"medods-test/internal/smtp"
	"medods-test/internal/tokens"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RefreshTokens(c echo.Context) error {
	cc := c.(*CustomContext)
	log := ctxvalue.GetLog(cc.Ctx)

	var refreshRequest TokenRequest
	if err := c.Bind(&refreshRequest); err != nil {
		log.Error("bind TokenRequest", "error", err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	refreshTokenHashed, err := dbhandlers.GetRefreshHash(cc.Ctx, refreshRequest.GUID)
	if err != nil {
		log.Error("get refresh hash", "error", err)
		return c.String(http.StatusInternalServerError, "get refresh hash")
	}

	if err := bcrypt.CompareHashAndPassword(refreshTokenHashed, []byte(refreshRequest.RefreshToken)); err != nil {
		log.Error("compare hash", "errror", err)
		return c.String(http.StatusForbidden, "invalid token")
	}

	log.Info("update", "token", refreshRequest.RefreshToken, "hash", string(refreshTokenHashed))

	ipFromDB, err := dbhandlers.CheckIP(cc.Ctx, refreshRequest.GUID)
	if err != nil {
		log.Error("check ip", "error", err)
		return c.String(http.StatusInternalServerError, "check ip")
	}

	if ipFromDB != refreshRequest.IPAddress {
		email, err := dbhandlers.GetEmail(cc.Ctx, refreshRequest.GUID)
		if err != nil {
			log.Error("get email", "error", err)
			return c.String(http.StatusInternalServerError, "failed to send warning")
		}

		if err := smtp.SendWarning(email, refreshRequest.IPAddress); err != nil {
			log.Error("send warning", "error", err)
			return c.String(http.StatusInternalServerError, "failed to send warning")
		}

		log.Info("different ip address")
		return c.String(http.StatusForbidden, "different IP address")
	}

	t, err := tokens.CreateAccessToken(refreshRequest.IPAddress, "secret")
	if err != nil {
		log.Error("create access token", "error", err)
		return c.String(http.StatusInternalServerError, "create token")
	}

	refreshToken, refreshTokenHashed, err := tokens.CreateRefreshToken("secret")
	if err != nil {
		log.Error("create refresh token", "error", err)
		return c.String(http.StatusInternalServerError, "create refresh token")
	}

	if err := dbhandlers.UpdateRefreshToken(cc.Ctx, refreshRequest.GUID, refreshRequest.IPAddress, refreshTokenHashed); err != nil {
		log.Error("update refresh token", "error", err)
		return c.String(http.StatusInternalServerError, "update refresh token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access token":  t,
		"refresh token": refreshToken,
	})
}

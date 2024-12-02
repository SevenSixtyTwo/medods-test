package apihandlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Ctx context.Context
}

type TokenRequest struct {
	GUID         uuid.UUID `json:"guid"`
	RefreshToken string    `json:"refresh_token"`
	IPAddress    string    `json:"ip_address"`
}

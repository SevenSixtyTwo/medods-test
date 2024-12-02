package main

import (
	"context"
	"errors"
	"fmt"
	apihandlers "medods-test/internal/api-handlers"
	ctxvalue "medods-test/internal/ctx-value"
	"medods-test/internal/db"
	"medods-test/internal/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	log := logger.SetupLogger()

	log.Info("statrting API")

	dbp, err := db.GetPostgresDb(ctx)
	if err != nil {
		log.Error("db postgres", "error", err)
		panic(fmt.Errorf("get db postgres %v", err))
	}

	ctx = context.WithValue(ctx, ctxvalue.ValueDbPostgres, dbp)
	ctx = context.WithValue(ctx, ctxvalue.ValueLog, log)

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &apihandlers.CustomContext{c, ctx}
			return next(cc)
		}
	})

	e.POST("/api/auth/token", apihandlers.GetTokens)
	e.POST("/api/auth/refresh", apihandlers.RefreshTokens)

	if err := e.Start(":3030"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("failed to start server", "error", err)
		panic(fmt.Errorf("failed to start server %v", err))
	}
}

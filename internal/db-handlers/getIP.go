package dbhandlers

import (
	"context"
	"fmt"
	ctxvalue "medods-test/internal/ctx-value"
)

func GetIP(ctx context.Context, refreshToken []byte) (string, error) {
	db := ctxvalue.GetDbPostgres(ctx)
	query := `SELECT ip_address FROM auth_service.users WHERE token_hash = $1`

	var ip string
	row := db.QueryRow(ctx, query, refreshToken)
	if err := row.Scan(&ip); err != nil {
		return "", fmt.Errorf("query ip %v", err)
	}

	return ip, nil
}

package dbhandlers

import (
	"context"
	"fmt"
	ctxvalue "medods-test/internal/ctx-value"

	"github.com/google/uuid"
)

func CheckIP(ctx context.Context, id uuid.UUID) (string, error) {
	db := ctxvalue.GetDbPostgres(ctx)
	query := `SELECT ip_address FROM auth_service.users WHERE id = $1`

	// TODO: delete refresh_token

	var ip string
	row := db.QueryRow(ctx, query, id)
	if err := row.Scan(&ip); err != nil {
		return "", fmt.Errorf("query ip %v", err)
	}

	return ip, nil
}

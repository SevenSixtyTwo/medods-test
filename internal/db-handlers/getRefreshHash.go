package dbhandlers

import (
	"context"
	"fmt"
	ctxvalue "medods-test/internal/ctx-value"

	"github.com/google/uuid"
)

func GetRefreshHash(ctx context.Context, id uuid.UUID) ([]byte, error) {
	db := ctxvalue.GetDbPostgres(ctx)
	query := `SELECT token_hash FROM auth_service.users WHERE id = $1`

	// TODO: delete refresh_token

	var hash []byte
	row := db.QueryRow(ctx, query, id)
	if err := row.Scan(&hash); err != nil {
		return nil, fmt.Errorf("query hash %v", err)
	}

	return hash, nil
}

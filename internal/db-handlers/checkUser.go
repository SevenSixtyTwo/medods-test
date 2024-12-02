package dbhandlers

import (
	"context"
	ctxvalue "medods-test/internal/ctx-value"

	"github.com/google/uuid"
)

func CheckUser(ctx context.Context, id uuid.UUID) bool {
	db := ctxvalue.GetDbPostgres(ctx)
	query := `SELECT EXISTS(SELECT 1 FROM auth_service.users WHERE id = $1)`

	isExist := false
	row := db.QueryRow(ctx, query, id)
	if err := row.Scan(&isExist); err != nil {
		return false
	}

	return isExist
}

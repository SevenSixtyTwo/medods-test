package dbhandlers

import (
	"context"
	"fmt"
	ctxvalue "medods-test/internal/ctx-value"

	"github.com/google/uuid"
)

func GetEmail(ctx context.Context, id uuid.UUID) (string, error) {
	db := ctxvalue.GetDbPostgres(ctx)
	query := `SELECT email FROM auth_service.users WHERE id = $1`

	var email string
	row := db.QueryRow(ctx, query, id)
	if err := row.Scan(&email); err != nil {
		return "", fmt.Errorf("query email %v", err)
	}

	return email, nil
}

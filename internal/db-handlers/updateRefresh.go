package dbhandlers

import (
	"context"
	"fmt"
	ctxvalue "medods-test/internal/ctx-value"

	"github.com/google/uuid"
)

func UpdateRefreshToken(ctx context.Context, id uuid.UUID, ip string, refreshHash []byte) error {
	db := ctxvalue.GetDbPostgres(ctx)
	log := ctxvalue.GetLog(ctx)

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error("begin transaction", "error", err)
		return fmt.Errorf("begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	log.Debug("started DEPOSIT transaction")

	query := `UPDATE auth_service.users SET token_hash = $1, ip_address = $2 WHERE id = $3;`

	result, err := tx.Exec(ctx, query, refreshHash, ip, id)
	if err != nil {
		return fmt.Errorf("update bank accounts: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("wrong uuid")
	}

	log.Debug("executed DEPOSIT update")

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %v", err)
	}

	log.Debug("commited DEPOSIT transaction")

	return nil
}

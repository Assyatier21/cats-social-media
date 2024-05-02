package pkg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type contextDBType string

var ContextTxValue contextDBType = "TX"

// ExtractDB is used by other repo to extract the trx from context
func ExtractTx(ctx context.Context) (*sql.Tx, error) {

	db, ok := ctx.Value(ContextTxValue).(*sql.Tx)
	if !ok {
		return nil, fmt.Errorf("TX is not found in context")
	}

	return db, nil
}

// WithTransactionDB used for common transaction handling
// all the context must use the same database session.
type WithTransactionDB interface {
	BeginTransaction(ctx context.Context) (context.Context, error)
	CommitTransaction(ctx context.Context) error
	RollbackTransaction(ctx context.Context) error
}

// WithTransaction is helper function that simplify the transaction execution handling
func WithTransaction(ctx context.Context, trx WithTransactionDB, trxFunc func(dbCtx context.Context) error) error {
	dbCtx, err := trx.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err = trx.RollbackTransaction(dbCtx)
			fmt.Println("With transaction Error occured", p)
		} else if err != nil {
			err = trx.RollbackTransaction(dbCtx)

		} else {
			err = trx.CommitTransaction(dbCtx)

		}
	}()

	err = trxFunc(dbCtx)
	return err
}

type SqlWithTransactionService struct {
	db *sqlx.DB
}

func NewSqlWithTransactionService(db *sqlx.DB) *SqlWithTransactionService {
	return &SqlWithTransactionService{
		db: db,
	}
}

func (r *SqlWithTransactionService) BeginTransaction(ctx context.Context) (context.Context, error) {
	dbTrx, _ := r.db.BeginTx(ctx, nil)
	trxCtx := context.WithValue(ctx, ContextTxValue, dbTrx)
	return trxCtx, nil
}

func (r *SqlWithTransactionService) CommitTransaction(ctx context.Context) error {
	db, err := ExtractTx(ctx)
	if err != nil {
		return err
	}
	return db.Commit()
}

func (r *SqlWithTransactionService) RollbackTransaction(ctx context.Context) error {

	db, err := ExtractTx(ctx)
	if err != nil {
		return err
	}

	return db.Rollback()
}

package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

// Provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	store := Store{
		Queries: New(db),
		db:      db,
	}
	return store
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rollbackErr := tx.Rollback() // Store rollback error separately
		if rollbackErr != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rollbackErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId int64
	ToAccountId   int64
	Amount        int64
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx function performs a money transfer from one account to the other.
// It creates a transfer record, add account entries, and update accounts balance within a db transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var res TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Create Transfer record
		res.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		// FromAccount entry
		res.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		// ToAccount's entry
		res.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    +arg.Amount,
		})

		if err != nil {
			return err
		}

		// TODO: update account's which has to be considered for deadlock prevention
		return nil
	})
	return res, err
}

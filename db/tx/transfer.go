package tx

import (
	"context"

	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
)

type TransferTxParams struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        string `json:"amount"`
}

type TransferTxResult struct {
	Transfer    sqlc.Transfer `json:"transfer"`
	FromAccount sqlc.Account  `json:"from_account"`
	ToAccount   sqlc.Account  `json:"to_account"`
	FromEntry   sqlc.Entry    `json:"from_entry"`
	ToEntry     sqlc.Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		transfer, err := q.CreateTransfer(ctx, sqlc.CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		fromEntry, err := q.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    "-" + arg.Amount,
		})
		if err != nil {
			return err
		}

		toEntry, err := q.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		updatedAccount1, err := q.AddAccountBalance(ctx, sqlc.AddAccountBalanceParams{
			Amount: "-" + arg.Amount,
			ID:     arg.FromAccountID,
		})
		if err != nil {
			return err
		}

		updatedAccount2, err := q.AddAccountBalance(ctx, sqlc.AddAccountBalanceParams{
			Amount: arg.Amount,
			ID:     arg.ToAccountID,
		})
		if err != nil {
			return err
		}

		result = TransferTxResult{
			Transfer:    transfer,
			FromEntry:   fromEntry,
			ToEntry:     toEntry,
			FromAccount: updatedAccount1,
			ToAccount:   updatedAccount2,
		}

		return nil
	})

	return result, err
}

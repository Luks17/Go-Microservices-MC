package repository

import (
	"context"

	"github.com/Luks17/Go-Microservices-MC/db/repository/sqlc"
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

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, sqlc.CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    "-" + arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, sqlc.CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		fromAccountParams := sqlc.AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: "-" + arg.Amount,
		}
		toAccountParams := sqlc.AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		}
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, fromAccountParams, toAccountParams)
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, toAccountParams, fromAccountParams)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *sqlc.Queries,
	addAccountBalanceParams1,
	addAccountBalanceParams2 sqlc.AddAccountBalanceParams,
) (account1, account2 sqlc.Account, err error) {
	account1, err = q.AddAccountBalance(ctx, addAccountBalanceParams1)
	if err != nil {
		return account1, account2, err
	}

	account2, err = q.AddAccountBalance(ctx, addAccountBalanceParams2)

	return account1, account2, err
}

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	TransferTx(ctx context.Context, args TransferArgs) (TransferResult, error)
	CreateAccount(ctx context.Context, args CreateAccountArgs) (Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetAccounts(ctx context.Context, args GetAccountsArgs) ([]Account, error)
	UpdateAccount(ctx context.Context, args UpdateAccountArgs) (Account, error)
	GetEntries(ctx context.Context, args GetEntriesArgs) ([]Entry, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	UpdateEntryAmount(ctx context.Context, args UpdateEntryArgs) (Entry, error)
	CreateTxn(ctx context.Context, args CreateTxnArgs) (Txn, error)
	DeleteTxn(ctx context.Context, id int64) error
	GetTxn(ctx context.Context, id int64) (Txn, error)
	GetTxns(ctx context.Context, args GetTxnArgs) ([]Txn, error)
	UpdateTxnAmount(ctx context.Context, args UpdateTxnArgs) (Txn, error)
	CreateEntry(ctx context.Context, args CreateEntryArgs) (Entry, error)
	DeleteEntry(ctx context.Context, id int64) error
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, cb func(*Queries) error) error {

	opts := &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	}
	tx, err := store.db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}
	q := New(tx)
	err = cb(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Transaction Error: %v, Rollback error:; %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferResult struct {
	Txn         Txn
	FromAccount Account
	ToAccount   Account
	FromEntry   Entry
	ToEntry     Entry
}
type TransferArgs struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
}

// create transaction, //create from entry, create to entry, change from account balance, change to account balance
func (store *SQLStore) TransferTx(ctx context.Context, args TransferArgs) (TransferResult, error) {
	var result TransferResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Txn, err = q.CreateTxn(ctx, CreateTxnArgs{From_account_id: args.FromAccountID, To_account_id: args.ToAccountID, Amount: args.Amount})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryArgs{Account_id: args.FromAccountID, Amount: -args.Amount})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryArgs{Account_id: args.ToAccountID, Amount: args.Amount})
		if err != nil {
			return err
		}
		// get the two accounts so you can update their balances
		result.FromAccount, err = q.GetAccountForUpdate(ctx, args.FromAccountID)
		if err != nil {
			return err
		}
		result.ToAccount, err = q.GetAccountForUpdate(ctx, args.ToAccountID)
		if err != nil {
			return err
		}
		// update their balances
		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountArgs{
			Balance:  result.FromAccount.Balance - args.Amount,
			Currency: result.FromAccount.Currency,
			Id:       args.FromAccountID,
		})
		if err != nil {
			return err
		}
		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountArgs{
			Balance:  result.ToAccount.Balance + args.Amount,
			Currency: result.ToAccount.Currency,
			Id:       args.ToAccountID,
		})
		return nil
	})
	return result, err
}

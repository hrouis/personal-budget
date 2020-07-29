package account

import (
	"context"
	"personal-budget/internal/data"
	"personal-budget/internal/database"
)

const (
	revenues string = "revenues"
	expenses        = "expenses"
	accounts        = "aacounts"
)

//Repository interface.
type Repository interface {
	Create(ctx context.Context, account data.Account) (string, error)
	Read(ctx context.Context, documentID string) (data.Account, error)
	Update(ctx context.Context, account data.Account)
}

//Repo definition
type Repo struct {
	Connection *database.Connection
}

//NewRepo creates a new Repository
func NewRepo(ctx context.Context, config database.Config) Repository {
	connection := database.NewConnection(ctx, config)
	return &Repo{
		Connection: connection,
	}
}

//Create add an account to the database
func (repository *Repo) Create(ctx context.Context, account data.Account) (string, error) {
	repository.Connection.CreateDocument(ctx, accounts, account)
	//TODO implement return
	return "success", nil
}

//GetAccount gets the account from the dtabase.
func (repository *Repo) Read(ctx context.Context, documentID string) (data.Account, error) {
	document := repository.Connection.ReadDocument(ctx, accounts, documentID)
	return document.(data.Account), nil
}

//Update the account and saves to the database.
func (repository *Repo) Update(ctx context.Context, account data.Account) {
	repository.Connection.UpdateDocument(ctx, accounts, account.AccountID, account)
}

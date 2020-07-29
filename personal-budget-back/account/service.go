package account

import (
	"context"
	"personal-budget/internal/data"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

//Service interface
type Service interface {
	CreateAccount(ctx context.Context, availableAmount int64) (string, error)
	AddRevenue(ctx context.Context, revenue data.Revenue, accountID string) (string, error)
}

//AccountService implementation
type service struct {
	repository Repository
	logger     log.Logger
}

//NewService Account service
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

//CreateAccount method implementation
func (s service) CreateAccount(ctx context.Context, availableAmount int64) (string, error) {
	logger := log.With(s.logger, "method", "CreateAccount")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	account := data.Account{
		AccountID:       id,
		AvailableAmount: availableAmount,
	}
	logger.Log("Create Account ", id)
	return s.repository.Create(ctx, account)
}

func (s service) AddRevenue(ctx context.Context, revenue data.Revenue, accountID string) (string, error) {
	logger := log.With(s.logger, "method", "AddRevenue")
	account, _ := s.repository.Read(ctx, accountID)
	account.Revenues = append(account.Revenues, &revenue)
	s.repository.Update(ctx, account)
	logger.Log("Added Revenue", revenue)
	return "success", nil
}

package account

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

type (
	createAccountRequest struct {
		AvailableAmount string `json:"amount"`
	}
	createAccountResponse struct {
		Status    string `json:"status"`
		AccountID string `json:"accountId"`
		Error     string `json:"error"`
	}
	addRevenueRequest struct {
		Label     string  `json:"label"`
		AccountID string  `json:"accountId"`
		Amount    float64 `json:"amount"`
		Frequency string  `json:"frequency"`
	}
	addRevenueResponse struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}
)

func makeCreateAccountEndPoint(s Service) endpoint.Endpoint {
	return func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(createAccountRequest)
		amount, err := strconv.ParseFloat(req.AvailableAmount, 64)
		accountID, err := s.CreateAccount(context, amount)
		return createAccountResponse{
			Status:    "status",
			AccountID: accountID,
			Error:     err.Error(),
		}, nil
	}
}

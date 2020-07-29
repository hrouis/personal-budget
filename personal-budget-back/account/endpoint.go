package account

type (
	CreateAccountRequest struct {
		AvailableAmount string `json: "amount"`
	}
	CreateAccountResponse struct {
		Status    string `json: "status"`
		AccountID string `json: "accountId"`
		Error     string `json: "error"`
	}
)

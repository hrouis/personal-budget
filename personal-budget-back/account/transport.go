package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type errorer interface {
	error() error
}

func decodeCreateAccountRequest(_ context.Context, request *http.Request) (interface{}, error) {
	vars := mux.Vars(request)
	amount, ok := vars["amount"]
	if !ok {
		return nil, nil
	}
	return createAccountRequest{
		AvailableAmount: amount,
	}, nil
}
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode error
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// MakeHandler method
func MakeHandler(ctx context.Context, accountService Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	accountHandler := kithttp.NewServer(
		makeCreateAccountEndPoint(accountService),
		decodeCreateAccountRequest,
		encodeResponse,
		options...,
	)
	r.Handle("/account/v1/create", accountHandler).Methods("POST")
	return r
}

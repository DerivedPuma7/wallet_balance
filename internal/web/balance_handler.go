package web

import (
	"encoding/json"
	"net/http"

	"github.com.br/derivedpuma7/balance/internal/usecase/get_balance_by_account"
	"github.com/go-chi/chi/v5"
)

type WebBalanceHandler struct {
  GetBalanceByAccountUseCase get_balance_by_account.GetBalanceByAccountUseCase
}

func NewWebBalanceHandler(getBalanceByAccountUseCase get_balance_by_account.GetBalanceByAccountUseCase) *WebBalanceHandler {
  return &WebBalanceHandler{
    GetBalanceByAccountUseCase: getBalanceByAccountUseCase,
  }
}

func (h *WebBalanceHandler) Handle(w http.ResponseWriter, r *http.Request) {
  accountId := chi.URLParam(r, "account_id")

  output, err := h.GetBalanceByAccountUseCase.Execute(get_balance_by_account.GetBalanceByAccountInput{
    AccountId: accountId,
  })
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  err = json.NewEncoder(w).Encode(output)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
}

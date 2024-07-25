package get_balance_by_account

import (
	"errors"

	"github.com.br/derivedpuma7/balance/internal/gateway"
)

type GetBalanceByAccountInput struct {
  AccountId string
}

type GetBalanceByAccountOutput struct {
  Balance float64
}

type GetBalanceByAccountUseCase struct {
  BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceByAccountUseCase(balanceGateway gateway.BalanceGateway) *GetBalanceByAccountUseCase {
  return &GetBalanceByAccountUseCase{
    BalanceGateway: balanceGateway,
  }
}

func (uc *GetBalanceByAccountUseCase) Execute(input GetBalanceByAccountInput) (*GetBalanceByAccountOutput, error) {
  accountBalance, err := uc.BalanceGateway.FindByAccountId(input.AccountId)
  if err != nil {
    return nil, err
  }
  if accountBalance == nil {
    return nil, errors.New("account does not exist")
  }

  return &GetBalanceByAccountOutput{
    Balance: accountBalance.Balance,
  }, nil
}

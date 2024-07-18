package update_balance

import (
	"github.com.br/derivedpuma7/balance/internal/entity"
	"github.com.br/derivedpuma7/balance/internal/gateway"
)

type UpdateBalanceInputDto struct {
  AccountId string
  Balance float64
}

type UpdateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewUpdateBalanceUseCase(balanceGateway gateway.BalanceGateway) *UpdateBalanceUseCase {
  return &UpdateBalanceUseCase{
    BalanceGateway: balanceGateway,
  }
}

func (uc *UpdateBalanceUseCase) Execute(input UpdateBalanceInputDto) error {
  existingBalance, err := uc.BalanceGateway.FindByAccountId(input.AccountId)
  if err != nil {
    return err
  }
  if existingBalance != nil {
    return uc.UpdateAccountBalance(existingBalance, input)
  }
  return uc.CreateAccountBalance(input)
}

func (uc *UpdateBalanceUseCase) UpdateAccountBalance(balance *entity.AccountBalance, input UpdateBalanceInputDto) error {
  err := balance.UpdateBalance(input.Balance)
  if err != nil {
    return err
  }
  err = uc.BalanceGateway.Update(balance)
  if err != nil {
    return err
  }
  return nil
}

func (uc *UpdateBalanceUseCase) CreateAccountBalance(input UpdateBalanceInputDto) error {
  accountBalance, err := entity.NewAccountBalance(input.AccountId, input.Balance)
  if err != nil {
    return err
  }
  err = uc.BalanceGateway.Save(accountBalance)
  if err != nil {
    return err
  }
  return nil
}

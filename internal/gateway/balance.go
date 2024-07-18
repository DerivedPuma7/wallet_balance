package gateway

import "github.com.br/derivedpuma7/balance/internal/entity"

type BalanceGateway interface {
  Save(balance *entity.AccountBalance) error
  Update(balance *entity.AccountBalance) error
  FindByAccountId(accountId string) (*entity.AccountBalance, error)
}

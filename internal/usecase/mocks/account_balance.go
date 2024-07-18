package mocks

import (
	"github.com.br/derivedpuma7/balance/internal/entity"
	"github.com.br/derivedpuma7/balance/internal/gateway"
	"github.com/stretchr/testify/mock"
)

type AccountBalanceMock struct {
	mock.Mock
}

var _ gateway.BalanceGateway = (*AccountBalanceMock)(nil)

func (a *AccountBalanceMock) FindByAccountId(accountId string) (*entity.AccountBalance, error) {
	args := a.Called(accountId)
	return args.Get(0).(*entity.AccountBalance), args.Error(1)
}

func (a *AccountBalanceMock) Save(balance *entity.AccountBalance) error {
	args := a.Called(balance)
	return args.Error(0)
}

func (a *AccountBalanceMock) Update(balance *entity.AccountBalance) error {
	args := a.Called(balance)
	return args.Error(0)
}

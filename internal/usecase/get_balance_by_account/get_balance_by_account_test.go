package get_balance_by_account

import (
	"errors"
	"testing"

	"github.com.br/derivedpuma7/balance/internal/entity"
	"github.com.br/derivedpuma7/balance/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBalanceByAccount_Execute(t *testing.T) {
  existingBalance, _ := entity.NewAccountBalance("id", 10)
  m := &mocks.AccountBalanceMock{}
  m.On("FindByAccountId", mock.Anything).Return(existingBalance, nil)
  uc := NewGetBalanceByAccountUseCase(m)

  result, err := uc.Execute(GetBalanceByAccountInput{
    AccountId: "any account id",
  })

  assert.Nil(t, err)
  assert.NotNil(t, result)
  assert.Equal(t, 10.0, result.Balance)
}

func TestGetBalanceByNonExistingAccount_Execute(t *testing.T) {
  m := &mocks.AccountBalanceMock{}
  m.On("FindByAccountId", mock.Anything).Return((*entity.AccountBalance)(nil), errors.New("any error"))
  uc := NewGetBalanceByAccountUseCase(m)

  result, err := uc.Execute(GetBalanceByAccountInput{
    AccountId: "any account id",
  })

  assert.NotNil(t, err)
  assert.Nil(t, result)
}

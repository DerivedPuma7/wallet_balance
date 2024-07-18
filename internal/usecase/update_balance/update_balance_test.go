package update_balance

import (
	"testing"

	"github.com.br/derivedpuma7/balance/internal/entity"
	"github.com.br/derivedpuma7/balance/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateBalanceNotExists_Execute(t *testing.T) {
  existingBalance := (*entity.AccountBalance)(nil)
  m := &mocks.AccountBalanceMock{}
  m.On("Save", mock.Anything).Return(nil)
  m.On("Update", mock.Anything).Return(nil)
  m.On("FindByAccountId", "any account id").Return(existingBalance, nil)
  uc := NewUpdateBalanceUseCase(m)

  err := uc.Execute(UpdateBalanceInputDto{
    AccountId: "any account id",
    Balance: 20,
  })

  assert.Nil(t, err)
  m.AssertNumberOfCalls(t, "Save", 1)
  m.AssertNotCalled(t, "Update")
}

func TestUpdateBalanceAlreadyExists_Execute(t *testing.T) {
  existingBalance, _ := entity.NewAccountBalance("id", 10)
  m := &mocks.AccountBalanceMock{}
  m.On("Save", mock.Anything).Return(nil)
  m.On("Update", mock.Anything).Return(nil)
  m.On("FindByAccountId", "any account id").Return(existingBalance, nil)
  uc := NewUpdateBalanceUseCase(m)

  err := uc.Execute(UpdateBalanceInputDto{
    AccountId: "any account id",
    Balance: 20,
  })

  assert.Nil(t, err)
  m.AssertNumberOfCalls(t, "Update", 1)
  m.AssertCalled(t, "Update", existingBalance)
  m.AssertNotCalled(t, "Save")
}

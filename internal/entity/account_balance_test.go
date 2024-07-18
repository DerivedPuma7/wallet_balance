package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewBalance(t *testing.T) {
  accountBalance, err := NewAccountBalance("any account id", 10)

  assert.Nil(t, err)
  assert.NotNil(t, accountBalance)
  assert.Equal(t, "any account id", accountBalance.AccountId)
  assert.Equal(t, 10.0, accountBalance.Balance)
}

func TestCreateNewBalanceWithInvalidBalance(t *testing.T) {
  accountBalance, err := NewAccountBalance("any account id", -10)

  assert.Nil(t, accountBalance)
  assert.NotNil(t, err)
  assert.Error(t, err, "balance cant be lower than zero")
}

func TestUpdateBalance(t *testing.T) {
  accountBalance, _ := NewAccountBalance("any account id", 10)

  err := accountBalance.UpdateBalance(20)

  assert.Nil(t, err)
  assert.Equal(t, 20.0, accountBalance.Balance)
}

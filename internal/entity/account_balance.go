package entity

import (
	"errors"
	"time"

  "github.com/google/uuid"
)

type AccountBalance struct {
  ID string
  AccountId string
  Balance float64
  UpdatedAt time.Time
}

func NewAccountBalance(accountId string, balance float64) (*AccountBalance, error) {
  newBalance := &AccountBalance{
    ID: uuid.New().String(),
    AccountId: accountId,
    Balance: balance,
    UpdatedAt: time.Now(),
  }
  err := newBalance.Validate()
  if err != nil {
    return nil, err
  }
  return newBalance, nil
}

func (b *AccountBalance) Validate() error {
  if b.Balance < 0 {
    return errors.New("balance cant be lower than zero")
  }
  return nil
}

func (b *AccountBalance) UpdateBalance(balance float64) error {
  if balance < 0 {
    return errors.New("balance cant be lower than zero")
  }
  b.Balance = balance
  return nil
}

package database

import (
	"database/sql"
	"time"

	"github.com.br/derivedpuma7/balance/internal/entity"
	"github.com.br/derivedpuma7/balance/internal/gateway"
)

type AccountBalanceDb struct {
	DB *sql.DB
}

func NewAccountBalanceDb(db *sql.DB) *AccountBalanceDb {
	return &AccountBalanceDb{
		DB: db,
	}
}

var _ gateway.BalanceGateway = (*AccountBalanceDb)(nil)

func (a *AccountBalanceDb) FindByAccountId(accountId string) (*entity.AccountBalance, error) {
  accountBalance := &entity.AccountBalance{}
	query := "SELECT id, accountId, balance, updatedAt FROM balances WHERE accountId = ?"
  stmt, err := a.DB.Prepare(query)
  if err != nil {
		return nil, err
	}
	defer stmt.Close()
  row := stmt.QueryRow(accountId)
  if err := row.Scan(&accountBalance.ID, &accountBalance.AccountId, &accountBalance.Balance, &accountBalance.UpdatedAt); err != nil {
    if err == sql.ErrNoRows {
      return nil, nil
    }
		return nil, err
	}
	return accountBalance, nil
}

func (a *AccountBalanceDb) Save(balance *entity.AccountBalance) error {
	sql := "INSERT INTO balances(id, accountId, balance, updatedAt) VALUES(?, ?, ?, ?)"
  stmt, err := a.DB.Prepare(sql)
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(balance.ID, balance.AccountId, balance.Balance, balance.UpdatedAt)
  if err != nil {
		return err
	}
	return nil
}

func (a *AccountBalanceDb) Update(balance *entity.AccountBalance) error {
	sql := "UPDATE balances SET balance = ?, updatedAt = ? WHERE id = ?"
  stmt, err := a.DB.Prepare(sql)
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, err = stmt.Exec(balance.Balance, time.Now(), balance.ID)
  if err != nil {
		return err
	}
	return nil
}

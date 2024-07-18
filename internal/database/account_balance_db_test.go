package database

import (
	"database/sql"
	"testing"

	"github.com.br/derivedpuma7/balance/internal/entity"
  _ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountBalanceDbTestSuite struct {
  suite.Suite
  db *sql.DB
  accountBalanceDb *AccountBalanceDb
}

func (a *AccountBalanceDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory")
	a.Nil(err)
	a.db = db
	db.Exec("CREATE TABLE balances (id varchar(255), accountId varchar(255), balance float, updatedAt date)")
  a.accountBalanceDb = NewAccountBalanceDb(db)
}

func(a *AccountBalanceDbTestSuite) TearDownSuite() {
	defer a.db.Close()
	a.db.Exec("DROP TABLE balances")
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(AccountBalanceDbTestSuite))
}

func (s *AccountBalanceDbTestSuite) TestSave() {
	accountBalance, _ := entity.NewAccountBalance("any id", 10)
	
  err := s.accountBalanceDb.Save(accountBalance)

	s.Nil(err)
}

func (s *AccountBalanceDbTestSuite) TestFindByAccountId() {
	accountBalance, _ := entity.NewAccountBalance("any id", 10)
	s.accountBalanceDb.Save(accountBalance)

  accountBalance, err := s.accountBalanceDb.FindByAccountId(accountBalance.AccountId)

	s.Nil(err)
  s.Equal("any id", accountBalance.AccountId)
  s.Equal(10.0, accountBalance.Balance)
}

func (s *AccountBalanceDbTestSuite) TestUpdate() {
	accountBalance, _ := entity.NewAccountBalance("any id", 10)
	s.accountBalanceDb.Save(accountBalance)

  accountBalance.UpdateBalance(20)
  err := s.accountBalanceDb.Update(accountBalance)

	s.Nil(err)
}

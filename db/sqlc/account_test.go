package db

import (
	"context"
	"testing"
  "github.com/ademhoskin/GoBank/util"
	"github.com/stretchr/testify/require"
  "time"
)

func CreateRandomAccount(t *testing.T) Account {
  arg := CreateAccountParams{
    Owner: util.RandomName(),
    Balance: util.RandomNumeric(),
  }

  account, err := testQueries.CreateAccount(context.Background(), arg)
  require.NoError(t, err)
  require.NotEmpty(t, account)

  require.Equal(t, arg.Owner, account.Owner)
  require.Equal(t, arg.Balance, account.Balance)

  require.NotZero(t, account.ID)
  require.NotZero(t, account.CreatedAt)

  return account
}

func TestCreateAccount(t *testing.T) {
  CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
  account1 := CreateRandomAccount(t)
  account2, err := testQueries.GetAccount(context.Background(), account1.ID)

  require.NoError(t, err)
  require.NotEmpty(t, account2)

  require.Equal(t, account1.Owner, account2.Owner)
  require.Equal(t, account1.Balance, account2.Balance)
  require.Equal(t, account1.ID, account2.ID)
  require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
  for i := 0; i < 10; i++ {
    CreateRandomAccount(t)
  }

  arg := ListAccountsParams{
    Limit: 5,
    Offset: 5,
  }

  list, err := testQueries.ListAccounts(context.Background(), arg)

  require.NoError(t, err)
  require.Len(t, list, 5)

  for _, account := range list {
    require.NotEmpty(t, account)
  }
}

func TestUpdateAccount(t *testing.T) {
  account1 := CreateRandomAccount(t)

  arg := UpdateAccountBalanceParams{
    ID: account1.ID,
    Balance: util.RandomNumeric(),
  }

  err:= testQueries.UpdateAccountBalance(context.Background(), arg)

  require.NoError(t, err)

  account2, err := testQueries.GetAccount(context.Background(), account1.ID)

  require.NoError(t, err)
  require.NotEmpty(t, account2)

  require.Equal(t, account1.Owner, account2.Owner)
  require.Equal(t, arg.Balance, account2.Balance)
  require.Equal(t, account1.ID, account2.ID)
  require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
  account1 := CreateRandomAccount(t)
  err := testQueries.DeleteAccount(context.Background(), account1.ID)
  require.NoError(t, err)

  account2, err2 := testQueries.GetAccount(context.Background(), account1.ID)
  require.Error(t, err2)
  require.Empty(t, account2)

}

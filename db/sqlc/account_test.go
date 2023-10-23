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
    Owner: util.RandomOwner(),
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
  


 

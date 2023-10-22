package db

import (
	"context"
	"testing"
  
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "John",
		Balance: "1000.00",
  }


	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	
  require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

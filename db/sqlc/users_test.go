package db

import (
	"context"
	//"log"
	"testing"

	"github.com/MeganViga/SimpleBookStore/util"
	"github.com/stretchr/testify/require"
)



func CreateRandomUser(t *testing.T)User{
	arg := CreateUserParams{
		Name: util.RandomUserName(),
		Email: util.RandomUserEmail(),
	}

	account, err := testQueries.CreateUser(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t, account)
	require.Equal(t,account.Name,arg.Name)
	require.Equal(t,account.Email,arg.Email)
	return account
}
func TestCreateUser(t *testing.T){
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T){
	user := CreateRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(),user.ID)
	require.NoError(t,err)
	require.NotEmpty(t,user2)
	require.Equal(t,user.ID,user2.ID)
	require.Equal(t,user.Name,user2.Name)
	require.Equal(t,user.Email,user2.Email)
	require.Equal(t,user.CreatedAt,user2.CreatedAt)

}

func TestUpdateEmail(t *testing.T){
	user := CreateRandomUser(t)
	arg := UpdateEmailParams{
		ID: user.ID,
		Email: util.RandomUserEmail(),
	}
	user2, err := testQueries.UpdateEmail(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,user2)
	require.Equal(t,user.ID,user2.ID)
	require.Equal(t,arg.Email,user2.Email)
	require.Equal(t,user.Name,user2.Name)
	require.NotEqual(t,user.Email,user2.Email)
}


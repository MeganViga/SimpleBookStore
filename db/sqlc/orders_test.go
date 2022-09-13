package db

import (
	"context"
	"testing"

	"github.com/MeganViga/SimpleBookStore/util"
	"github.com/stretchr/testify/require"
)
func CreateRandomOrder(t *testing.T)Order{
	book := CreateRandomBook(t)
	user := CreateRandomUser(t)
	quantity := util.RandomQuantity()
	arg := CreateOrderParams{
		BookID: book.ID,
		UserID: user.ID,
		Quantity: quantity,
		TotalPrice: float64(quantity) * book.Price,
		Status: true,
	}
	order, err := testQueries.CreateOrder(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,order)
	require.Equal(t,book.ID,order.BookID)
	require.Equal(t,user.ID,order.UserID)
	require.Equal(t,quantity,order.Quantity)
	require.Equal(t,float64(quantity)*book.Price,order.TotalPrice)
	require.Equal(t,arg.Status,order.Status)
	return order
}
func TestCreateOrder(t *testing.T){
	CreateRandomOrder(t)

}

func TestCancelOrder(t *testing.T){
	order := CreateRandomOrder(t)

	cancelledOrder, err := testQueries.CancelOrder(context.Background(),order.ID)
	require.NoError(t,err)
	require.NotEmpty(t,cancelledOrder)
	require.Equal(t,order.ID,cancelledOrder.ID)
	require.Equal(t,order.BookID,cancelledOrder.BookID)
	require.Equal(t,order.UserID,cancelledOrder.UserID)
	require.Equal(t,order.Quantity,cancelledOrder.Quantity)
	require.Equal(t,order.TotalPrice,cancelledOrder.TotalPrice)
	require.NotEqual(t,order.Status,cancelledOrder.Status)
}

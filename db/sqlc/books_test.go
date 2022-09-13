package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/MeganViga/SimpleBookStore/util"
)



func CreateRandomBook(t *testing.T)Book{
	arg := CreateBookParams{
		Title: util.RandomBookName(),
		Price: util.RandomPrice(),
		Stock:10,
	}
	book, err := testQueries.CreateBook(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,book)
	require.Equal(t,arg.Title,book.Title)
	require.Equal(t,arg.Price,book.Price)
	require.Equal(t,arg.Stock,book.Stock)
	return book
}
func TestCreateBooks(t *testing.T){
	CreateRandomBook(t)
}

func TestGetBook(t *testing.T){
	book := CreateRandomBook(t)
	book2 , err := testQueries.GetBook(context.Background(),book.ID)
	require.NoError(t,err)
	require.NotEmpty(t,book2)
	require.Equal(t,book.ID,book2.ID)
	require.Equal(t,book.Title,book2.Title)
	require.Equal(t,book.Stock,book2.Stock)
	require.Equal(t,book.CreatedAt,book2.CreatedAt)
	require.Equal(t,book.Price,book2.Price)
}

func TestUpdateBookStock(t *testing.T){
	book := CreateRandomBook(t)
	arg := UpdateBookStockParams{
		ID: book.ID,
		Stock: 12,
	}
	updatedbook, err := testQueries.UpdateBookStock(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,updatedbook)
	require.Equal(t,book.ID,updatedbook.ID)
	require.Equal(t,book.Price,updatedbook.Price)
	require.Equal(t,book.Title,updatedbook.Title)
	require.NotEqual(t,book.Stock,updatedbook.Stock)

}

func TestUpdateBookPrice(t *testing.T){
	book := CreateRandomBook(t)
	arg := UpdateBookPriceParams{
		ID: book.ID,
		Price: util.RandomPrice(),
	}
	updatedbook, err := testQueries.UpdateBookPrice(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,updatedbook)
	require.Equal(t,book.ID,updatedbook.ID)
	require.Equal(t,book.Title,updatedbook.Title)
	require.Equal(t,book.Stock,updatedbook.Stock)
	require.NotEqual(t,book.Price,updatedbook.Price)
	require.Equal(t,arg.Price,updatedbook.Price)
}

func TestDeleteBook(t *testing.T){
	book := CreateRandomBook(t)
	id, err := testQueries.DeleteBook(context.Background(),book.ID)
	require.NoError(t,err)
	require.NotEmpty(t,id)
	require.Equal(t,book.ID,id)
	book2, err := testQueries.GetBook(context.Background(),id)
	require.Error(t,err)
	require.Empty(t,book2)
}
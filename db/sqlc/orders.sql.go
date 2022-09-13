// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: orders.sql

package db

import (
	"context"
)

const cancelOrder = `-- name: CancelOrder :one
UPDATE orders SET status=false WHERE id = $1 RETURNING id, book_id, user_id, quantity, total_price, status, created_at
`

func (q *Queries) CancelOrder(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, cancelOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.UserID,
		&i.Quantity,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  book_id, user_id, quantity,total_price,status
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, book_id, user_id, quantity, total_price, status, created_at
`

type CreateOrderParams struct {
	BookID     int64   `json:"book_id"`
	UserID     int64   `json:"user_id"`
	Quantity   int32   `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status     bool    `json:"status"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.BookID,
		arg.UserID,
		arg.Quantity,
		arg.TotalPrice,
		arg.Status,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.UserID,
		&i.Quantity,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

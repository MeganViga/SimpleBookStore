-- name: CreateOrder :one
INSERT INTO orders (
  book_id, user_id, quantity,total_price,status
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CancelOrder :one
UPDATE orders SET status=false WHERE id = $1 RETURNING *;
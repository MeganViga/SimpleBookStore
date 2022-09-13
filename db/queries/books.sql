-- name: CreateBook :one
INSERT INTO books (
  title, price, stock
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetBook :one
SELECT * FROM books WHERE id = $1 LIMIT 1;
-- name: UpdateBookStock :one
UPDATE books SET stock = $2 WHERE id = $1 RETURNING *;

-- name: UpdateBookPrice :one
UPDATE books SET price = $2 WHERE id = $1 RETURNING *;

-- name: DeleteBook :one
DELETE FROM books WHERE id = $1 RETURNING id;
-- name: CreateUser :one
INSERT INTO users (
  name, email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateEmail :one
UPDATE users SET email = $2 WHERE id = $1 RETURNING *;



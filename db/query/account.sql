-- name: createAccount :one
INSERT INTO accounts (
  owner, 
  balance
) VALUES (
$1, $2
) RETURNING *;

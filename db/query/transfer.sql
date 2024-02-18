-- name: CreateTransfer :one
INSERT INTO transfers (
from_account_id,
to_account_id,
amount,
created_at
) VALUES (
  $1, $2, $3, $4
)RETURNING *;


-- name: GetTransfer :one
SELECT * FROM accounts
WHERE id = $1 
LIMIT 1
OFFSET $2;

-- name: ListTransfer :many
SELECT * FROM accounts
ORDER BY id;
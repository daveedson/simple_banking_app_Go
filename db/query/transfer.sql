-- name: CreateTransfer :one
INSERT INTO transfers (
from_account_id,
to_account_id,
amount,
created_at
) VALUES (
  $1, $2, $3, $4
)RETURNING *;
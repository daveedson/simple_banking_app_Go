-- name: CreatEntry :one
INSERT INTO accounts (
 owner,
 balance,
 currency,
 created_at
) VALUES (
  $1, $2, $3, $4
)RETURNING *;


-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 
LIMIT 1;



-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;
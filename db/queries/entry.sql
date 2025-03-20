-- name: CreateEntry :one
INSERT INTO entries(
    account_id,
    amount
)
VALUES($1,$2)
RETURNING *;


-- name: GetAccountEntries :many
SELECT * FROM entries WHERE account_id=$1;

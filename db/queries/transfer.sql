-- name: CreateTransfer :one
INSERT INTO transfers(
    from_account_id,
    to_account_id,
    amount
)VALUES($1,$2,$3)RETURNING *;


-- name: GetTransfersByAccount :many
SELECT * FROM transfers WHERE from_account_id = $1;
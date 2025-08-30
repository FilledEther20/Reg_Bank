-- name: CreateUser :one
INSERT INTO users(
    username,
    hashed_password,
    full_name,
    email,
    password_changed_at
)VALUES(
    $1,$2,$3,$4,$5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username=$1 LIMIT 1;

-- name: UpdatePassword :one 
UPDATE users
SET hashed_password=$2
WHERE username=$1
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users WHERE username=$1;
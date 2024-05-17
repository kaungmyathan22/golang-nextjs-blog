-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (
  name, password, email
) VALUES (
    $1, $2, $3
);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, email, createdAt, updatedAt
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

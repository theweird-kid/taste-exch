-- name: CreateUser :one
INSERT INTO users (email, name, password, profile_url)
VALUES ($1, $2, $3, $4)
RETURNING id, name, email, profile_url, created_at;

-- name: GetUserByEmail :one
SELECT id, name, email, password, profile_url, created_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, name, email, profile_url, created_at
FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET name = COALESCE($1, email),
    profile_url = COALESCE($2, profile_url)
WHERE id = $3;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

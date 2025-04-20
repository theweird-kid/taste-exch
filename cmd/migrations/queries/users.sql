-- name: CreateUser :one
INSERT INTO users (email, name, description, password, profile_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, email, description, profile_url, created_at;

-- name: GetUserByEmail :one
SELECT id, name, email, description, password, profile_url, created_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, name, email, description, profile_url, created_at
FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET name = COALESCE($1, name),
    description = COALESCE($2, description),
    profile_url = COALESCE($3, profile_url)
WHERE id = $4;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

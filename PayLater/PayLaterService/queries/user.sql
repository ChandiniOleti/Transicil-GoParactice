-- name: CreateUser :execresult
INSERT INTO users (
    name,
    email,
    password   
)
VALUES (
    ?,
    ?,
    ?
);

-- name: GetUsers :many
SELECT *
FROM users;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?;

-- name: UpdateUser :exec
UPDATE users
SET
    name = ?,
    email = ?,
    password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: UpdateCurrentDue :exec
UPDATE users
SET
    current_due = ?
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;
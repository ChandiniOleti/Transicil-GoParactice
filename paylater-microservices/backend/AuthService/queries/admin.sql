-- name: CreateAdmin :execresult
INSERT INTO admins (
    name,
    email,
    password
)
VALUES (
    ?,
    ?,
    ?
);

-- name: GetAdminByEmail :one
SELECT *
FROM admins
WHERE email = ?;

-- name: GetAdminByID :one
SELECT *
FROM admins
WHERE id = ?;

-- name: GetAdmins :many
SELECT *
FROM admins;

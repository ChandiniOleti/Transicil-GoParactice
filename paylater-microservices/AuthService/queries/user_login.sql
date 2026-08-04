-- Temporary until User Service (Phase 2) exposes verify API.
-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;

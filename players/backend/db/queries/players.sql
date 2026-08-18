-- name: GetPlayers :many
SELECT *
FROM players
LIMIT ? OFFSET ?;

-- name: GetPlayerByID :one
SELECT *
FROM players
WHERE playerID = ?;

-- name: GetPlayersCount :one
SELECT COUNT(*) FROM players;
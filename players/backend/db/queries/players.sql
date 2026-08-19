-- name: GetPlayers :many
SELECT *
FROM players
LIMIT ? OFFSET ?;

-- name: GetPlayerByID :one
SELECT *
FROM players
WHERE playerID = ?;

-- name: GetPlayersCount :one
SELECT COUNT(*)
FROM players;

-- name: SearchPlayers :many
SELECT *
FROM players
WHERE LOWER(nameFirst) LIKE CONCAT('%', LOWER(?), '%')
   OR LOWER(nameLast) LIKE CONCAT('%', LOWER(?), '%')
   OR LOWER(CONCAT(nameFirst, ' ', nameLast))
      LIKE CONCAT('%', LOWER(?), '%')
   OR LOWER(CONCAT(nameFirst, nameLast))
      LIKE CONCAT('%', LOWER(REPLACE(?, ' ', '')), '%')
LIMIT ? OFFSET ?;

-- name: SearchPlayersCount :one
SELECT COUNT(*)
FROM players
WHERE LOWER(nameFirst) LIKE CONCAT('%', LOWER(?), '%')
   OR LOWER(nameLast) LIKE CONCAT('%', LOWER(?), '%')
   OR LOWER(CONCAT(nameFirst, ' ', nameLast))
      LIKE CONCAT('%', LOWER(?), '%')
   OR LOWER(CONCAT(nameFirst, nameLast))
      LIKE CONCAT('%', LOWER(REPLACE(?, ' ', '')), '%');
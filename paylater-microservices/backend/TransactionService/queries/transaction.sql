-- name: CreateTransaction :execresult
INSERT INTO transactions (
    user_id,
    merchant_id,
    amount,
    commission_percentage,
    commission_amount,
    transaction_type
)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
);

-- name: GetTransactions :many
SELECT *
FROM transactions;

-- name: GetTransactionByID :one
SELECT *
FROM transactions
WHERE id = ?;

-- name: GetTransactionsByUser :many
SELECT *
FROM transactions
WHERE user_id = ?;

-- name: GetTransactionsByMerchant :many
SELECT *
FROM transactions
WHERE merchant_id = ?;

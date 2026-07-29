-- name: GetMerchantFeeCollected :one
SELECT
    merchant_id,
    SUM(commission_amount) AS total_fee_collected
FROM transactions
WHERE merchant_id = ?
  AND transaction_type = 'PURCHASE'
GROUP BY merchant_id;

-- name: GetUsersWithDue :many
SELECT *
FROM users
WHERE current_due > 0
ORDER BY current_due DESC;

-- name: GetUserDue :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetUsersReachedCreditLimit :many
SELECT *
FROM users
WHERE current_due >= credit_limit
ORDER BY current_due DESC;

-- name: GetTotalUserDues :one
SELECT
    SUM(current_due) AS total_due
FROM users;
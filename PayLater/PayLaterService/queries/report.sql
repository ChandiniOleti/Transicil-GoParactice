-- name: GetAllUserDues :many
SELECT
    id,
    name,
    email,
    credit_limit,
    current_due
FROM users;

-- name: GetMerchantCommissionReport :many
SELECT
    m.id,
    m.merchant_name,
    SUM(t.commission_amount) AS total_commission
FROM merchants m
JOIN transactions t
ON m.id = t.merchant_id
GROUP BY m.id, m.merchant_name;

-- name: GetTotalOutstandingDue :one
SELECT
    SUM(current_due) AS total_due
FROM users;

-- name: GetTransactionReport :many
SELECT
    t.id,
    u.name AS user_name,
    m.merchant_name,
    t.amount,
    t.commission_percentage,
    t.commission_amount,
    t.transaction_date
FROM transactions t
JOIN users u
ON t.user_id = u.id
JOIN merchants m
ON t.merchant_id = m.id;

-- name: GetMerchantTransactionCount :many
SELECT
    m.id,
    m.merchant_name,
    COUNT(t.id) AS total_transactions
FROM merchants m
LEFT JOIN transactions t
ON m.id = t.merchant_id
GROUP BY m.id, m.merchant_name;
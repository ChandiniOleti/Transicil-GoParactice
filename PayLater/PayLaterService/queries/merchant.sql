-- name: CreateMerchant :execresult
INSERT INTO merchants (
    merchant_name,
    email,
    commission
) VALUES (
    ?,
    ?,
    ?
);

-- name: GetMerchants :many
SELECT *
FROM merchants;

-- name: GetMerchantByID :one
SELECT *
FROM merchants
WHERE id = ?;

-- name: UpdateMerchant :exec
UPDATE merchants
SET
    merchant_name = ?,
    email = ?,
    commission = ?
WHERE id = ?;

-- name: UpdateCommission :exec
UPDATE merchants
SET
    commission = ?
WHERE id = ?;

-- name: DeleteMerchant :exec
DELETE FROM merchants
WHERE id = ?;
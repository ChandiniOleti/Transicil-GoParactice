-- Temporary until Merchant Service (Phase 3) exposes verify API.
-- name: GetMerchantByEmail :one
SELECT *
FROM merchants
WHERE email = ?;

/**
 * Transaction API contracts (via API Gateway).
 *
 * Gateway endpoints:
 * - POST /transactions                         (JWT)
 * - GET  /transactions                         (JWT)
 * - GET  /transactions/:id                     (JWT)
 * - GET  /transactions/user/:user_id           (JWT)
 * - GET  /transactions/merchant/:merchant_id   (JWT)
 *
 * Decimal money fields are serialized as strings.
 * transaction_type ENUM('PURCHASE','PAYBACK').
 */

/** Matches transactions.transaction_type ENUM. */
export type TransactionType = 'PURCHASE' | 'PAYBACK'

/**
 * Transaction history DTO (TransactionResponse).
 * Nullable SQL fields are omitted when unset (Go omitempty on pointers).
 * transaction_date is RFC3339 UTC when present.
 */
export interface Transaction {
  id: number
  user_id: number
  merchant_id?: number
  amount: string
  commission_percentage: string
  commission_amount: string
  transaction_type: TransactionType
  transaction_date?: string
}

export type TransactionResponse = Transaction

/** List endpoints return a JSON array (never null; service uses make(..., 0, n)). */
export type TransactionListResponse = Transaction[]

/**
 * POST /transactions request body (handlers.TransactionRequest).
 * user_id comes from the JWT, not the body.
 */
export interface CreateTransactionRequest {
  merchant_id: number
  amount: string
}

/**
 * POST /transactions success body (201).
 * Shape differs from TransactionResponse — uses transaction_amount, etc.
 */
export interface CreateTransactionResponse {
  message: string
  user_id: number
  merchant_id: number
  transaction_type: TransactionType
  transaction_amount: string
  commission_percentage: string
  commission_amount: string
  updated_current_due: string
  available_credit: string
}

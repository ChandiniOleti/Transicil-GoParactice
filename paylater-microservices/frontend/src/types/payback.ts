/**
 * Payback API contracts (via API Gateway).
 *
 * Gateway endpoints:
 * - POST /payback (JWT; USER for own account, or ADMIN)
 */

/** POST /payback request body (PaybackRequest). */
export interface PaybackRequest {
  user_id: number
  amount: string
}

/**
 * POST /payback success body (200).
 * Distinct from CreateTransactionResponse / TransactionResponse.
 */
export interface PaybackResponse {
  message: string
  user_id: number
  amount_paid: string
  updated_current_due: string
  available_credit: string
}

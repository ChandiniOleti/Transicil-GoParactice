/**
 * Report API contracts (via API Gateway).
 *
 * Gateway endpoints (JWT):
 * - GET /reports/merchant-fee/:id
 * - GET /reports/users-with-due
 * - GET /reports/user-due/:id
 * - GET /reports/credit-limit-users
 * - GET /reports/total-dues
 *
 * User-shaped report rows match UserService UserResponse (no password).
 */

/** User row used by due / credit-limit reports (clients.UserReport). */
export interface UserReport {
  id: number
  name: string
  email: string
  credit_limit: string
  current_due: string
}

/** GET /reports/merchant-fee/:id */
export interface MerchantFeeResponse {
  merchant_id: number
  total_fee_collected: string
}

/** GET /reports/total-dues */
export interface TotalDuesResponse {
  total_due: string
}

/**
 * GET /reports/users-with-due
 * Always a JSON array (service uses make(..., 0)).
 */
export type UsersWithDueResponse = UserReport[]

/** GET /reports/user-due/:id — single user report row. */
export type UserDueResponse = UserReport

/**
 * GET /reports/credit-limit-users
 * Always a JSON array (service uses make(..., 0)).
 */
export type CreditLimitUsersResponse = UserReport[]

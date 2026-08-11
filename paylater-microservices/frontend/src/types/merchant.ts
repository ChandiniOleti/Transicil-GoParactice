/**
 * Merchant API contracts (via API Gateway).
 *
 * Gateway endpoints:
 * - POST   /merchants
 * - GET    /merchants              (JWT)
 * - GET    /merchants/:id          (JWT)
 * - PUT    /merchants/:id          (JWT)
 * - PATCH  /merchants/:id/commission (JWT)
 * - DELETE /merchants/:id          (JWT)
 *
 * commission is MySQL DECIMAL serialized as a string.
 */

import type { MessageResponse } from './auth'

/** Merchant DTO returned by MerchantService (password excluded). */
export interface Merchant {
  id: number
  merchant_name: string
  email: string
  phone: string
  commission: string
}

/** Alias matching backend MerchantResponse naming. */
export type MerchantResponse = Merchant

/**
 * GET /merchants always returns a JSON array
 * (service pre-allocates with make(..., 0, n)).
 */
export type MerchantListResponse = Merchant[]

/** POST /merchants request body (CreateMerchantParams). */
export interface CreateMerchantRequest {
  merchant_name: string
  phone: string
  commission: string
  email: string
  password: string
}

/** POST /merchants success body. */
export type CreateMerchantResponse = MessageResponse

/**
 * PUT /merchants/:id request body.
 * Path param supplies id; handler overwrites any body id after binding.
 */
export interface UpdateMerchantRequest {
  merchant_name: string
  phone: string
  commission: string
  email: string
  password: string
}

/** PUT /merchants/:id success body. */
export type UpdateMerchantResponse = MessageResponse

/**
 * PATCH /merchants/:id/commission request body.
 * Path param supplies id.
 */
export interface UpdateCommissionRequest {
  commission: string
}

/** PATCH /merchants/:id/commission success body. */
export type UpdateCommissionResponse = MessageResponse

/** DELETE /merchants/:id success body. */
export type DeleteMerchantResponse = MessageResponse

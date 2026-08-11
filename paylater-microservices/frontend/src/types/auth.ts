/**
 * Auth / shared API contracts (via API Gateway).
 *
 * Gateway endpoints:
 * - POST /login
 * - POST /admin/login
 * - POST /merchant/login
 * - POST /admins (JWT)
 * - GET  /admins (JWT)
 *
 * Backend error envelope (shared/response + handlers): { "error": string }
 * Backend message envelope: { "message": string }
 */

/** Roles embedded in JWT claims by AuthService. */
export type AuthRole = 'USER' | 'ADMIN' | 'MERCHANT'

/** Consistent error body returned by handlers and gateway JWT middleware. */
export interface ApiErrorResponse {
  error: string
}

/** Consistent success message body used by many create/update/delete handlers. */
export interface MessageResponse {
  message: string
}

/** Shared login body for user, admin, and merchant login. */
export interface LoginRequest {
  email: string
  password: string
}

/**
 * Login success body. The `message` text differs by endpoint:
 * - User:     "User Login Successful"
 * - Admin:    "Admin Login Successful"
 * - Merchant: "Merchant Login Successful"
 */
export interface LoginResponse {
  message: string
  token: string
}

export type UserLoginRequest = LoginRequest
export type UserLoginResponse = LoginResponse

export type AdminLoginRequest = LoginRequest
export type AdminLoginResponse = LoginResponse

export type MerchantLoginRequest = LoginRequest
export type MerchantLoginResponse = LoginResponse

/** POST /admins request body (CreateAdminParams). */
export interface CreateAdminRequest {
  name: string
  email: string
  password: string
}

/** POST /admins success body. */
export interface CreateAdminResponse {
  message: string
}

/**
 * Admin row returned by GET /admins.
 * Note: AuthService returns generated.Admin including the password hash.
 */
export interface Admin {
  id: number
  name: string
  email: string
  password: string
}

/**
 * GET /admins returns []Admin.
 * When empty, Go's nil slice encodes as JSON null.
 */
export type AdminListResponse = Admin[] | null

/**
 * Claims stored in AuthService JWTs (jwt.MapClaims).
 * `exp` is a Unix timestamp (seconds).
 */
export interface JwtClaims {
  user_id: number
  email: string
  role: AuthRole
  exp: number
}

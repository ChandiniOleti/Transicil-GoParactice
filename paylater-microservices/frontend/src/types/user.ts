/**
 * User API contracts (via API Gateway).
 *
 * Gateway endpoints:
 * - POST   /users          (public registration)
 * - GET    /users          (JWT)
 * - GET    /users/:id      (JWT)
 * - PUT    /users/:id      (JWT)
 * - DELETE /users/:id      (JWT)
 *
 * Money fields (credit_limit, current_due) are MySQL DECIMAL values
 * serialized by the Go services as strings.
 */

import type { MessageResponse } from './auth'

/** User DTO returned by UserService (password excluded). */
export interface User {
  id: number
  name: string
  email: string
  credit_limit: string
  current_due: string
}

/** Alias matching backend UserResponse naming. */
export type UserResponse = User

/**
 * GET /users returns []UserResponse.
 * When the list is empty, Go's nil slice encodes as JSON null.
 */
export type UserListResponse = User[] | null

/** POST /users request body (CreateUserParams). credit_limit/current_due use DB defaults. */
export interface CreateUserRequest {
  name: string
  email: string
  password: string
}

/** POST /users success body. */
export type CreateUserResponse = MessageResponse

/**
 * PUT /users/:id request body.
 * Path param supplies id; handler overwrites any body id after binding.
 */
export interface UpdateUserRequest {
  name: string
  email: string
  password: string
}

/** PUT /users/:id success body. */
export type UpdateUserResponse = MessageResponse

/** DELETE /users/:id success body. */
export type DeleteUserResponse = MessageResponse

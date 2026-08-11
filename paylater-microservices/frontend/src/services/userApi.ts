import { api } from './api'
import type {
  CreateUserRequest,
  CreateUserResponse,
  DeleteUserResponse,
  UpdateUserRequest,
  UpdateUserResponse,
  UserListResponse,
  UserResponse,
} from '../types/user'

/** POST /users */
export async function createUser(
  payload: CreateUserRequest,
): Promise<CreateUserResponse> {
  const response = await api.post<CreateUserResponse>('/users', payload)
  return response.data
}

/** GET /users */
export async function getUsers(): Promise<UserListResponse> {
  const response = await api.get<UserListResponse>('/users')
  return response.data
}

/** GET /users/:id */
export async function getUserById(userId: number): Promise<UserResponse> {
  const response = await api.get<UserResponse>(`/users/${userId}`)
  return response.data
}

/** PUT /users/:id */
export async function updateUser(
  userId: number,
  payload: UpdateUserRequest,
): Promise<UpdateUserResponse> {
  const response = await api.put<UpdateUserResponse>(
    `/users/${userId}`,
    payload,
  )
  return response.data
}

/** DELETE /users/:id */
export async function deleteUser(userId: number): Promise<DeleteUserResponse> {
  const response = await api.delete<DeleteUserResponse>(`/users/${userId}`)
  return response.data
}

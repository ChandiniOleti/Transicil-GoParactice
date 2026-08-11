import { api } from './api'
import type {
  AdminListResponse,
  AdminLoginRequest,
  AdminLoginResponse,
  CreateAdminRequest,
  CreateAdminResponse,
  MerchantLoginRequest,
  MerchantLoginResponse,
  UserLoginRequest,
  UserLoginResponse,
} from '../types/auth'

/** POST /login */
export async function login(
  payload: UserLoginRequest,
): Promise<UserLoginResponse> {
  const response = await api.post<UserLoginResponse>('/login', payload)
  return response.data
}

/** POST /admin/login */
export async function adminLogin(
  payload: AdminLoginRequest,
): Promise<AdminLoginResponse> {
  const response = await api.post<AdminLoginResponse>('/admin/login', payload)
  return response.data
}

/** POST /merchant/login */
export async function merchantLogin(
  payload: MerchantLoginRequest,
): Promise<MerchantLoginResponse> {
  const response = await api.post<MerchantLoginResponse>(
    '/merchant/login',
    payload,
  )
  return response.data
}

/** POST /admins */
export async function createAdmin(
  payload: CreateAdminRequest,
): Promise<CreateAdminResponse> {
  const response = await api.post<CreateAdminResponse>('/admins', payload)
  return response.data
}

/** GET /admins */
export async function getAdmins(): Promise<AdminListResponse> {
  const response = await api.get<AdminListResponse>('/admins')
  return response.data
}

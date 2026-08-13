import { api } from './api'
import { buildEncryptedLoginBody } from '../utils/loginCrypto'
import type {
  AdminListResponse,
  AdminLoginResponse,
  CreateAdminRequest,
  CreateAdminResponse,
  MerchantLoginResponse,
  UserLoginResponse,
} from '../types/auth'

/** POST /login */
export async function login(
  email: string,
  password: string,
): Promise<UserLoginResponse> {
  const body = await buildEncryptedLoginBody(email, password)
  const response = await api.post<UserLoginResponse>('/login', body)
  return response.data
}

/** POST /admin/login */
export async function adminLogin(
  email: string,
  password: string,
): Promise<AdminLoginResponse> {
  const body = await buildEncryptedLoginBody(email, password)
  const response = await api.post<AdminLoginResponse>('/admin/login', body)
  return response.data
}

/** POST /merchant/login */
export async function merchantLogin(
  email: string,
  password: string,
): Promise<MerchantLoginResponse> {
  const body = await buildEncryptedLoginBody(email, password)
  const response = await api.post<MerchantLoginResponse>(
    '/merchant/login',
    body,
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

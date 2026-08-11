import { api } from './api'
import type {
  CreditLimitUsersResponse,
  MerchantFeeResponse,
  TotalDuesResponse,
  UserDueResponse,
  UsersWithDueResponse,
} from '../types/report'

/** GET /reports/merchant-fee/:id */
export async function getMerchantFeeReport(
  merchantId: number,
): Promise<MerchantFeeResponse> {
  const response = await api.get<MerchantFeeResponse>(
    `/reports/merchant-fee/${merchantId}`,
  )
  return response.data
}

/** GET /reports/users-with-due */
export async function getUsersWithDueReport(): Promise<UsersWithDueResponse> {
  const response = await api.get<UsersWithDueResponse>(
    '/reports/users-with-due',
  )
  return response.data
}

/** GET /reports/user-due/:id */
export async function getUserDueReport(
  userId: number,
): Promise<UserDueResponse> {
  const response = await api.get<UserDueResponse>(
    `/reports/user-due/${userId}`,
  )
  return response.data
}

/** GET /reports/credit-limit-users */
export async function getCreditLimitUsersReport(): Promise<CreditLimitUsersResponse> {
  const response = await api.get<CreditLimitUsersResponse>(
    '/reports/credit-limit-users',
  )
  return response.data
}

/** GET /reports/total-dues */
export async function getTotalDuesReport(): Promise<TotalDuesResponse> {
  const response = await api.get<TotalDuesResponse>('/reports/total-dues')
  return response.data
}

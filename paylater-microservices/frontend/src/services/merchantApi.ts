import { api } from './api'
import type {
  CreateMerchantRequest,
  CreateMerchantResponse,
  DeleteMerchantResponse,
  MerchantListResponse,
  MerchantResponse,
  UpdateCommissionRequest,
  UpdateCommissionResponse,
  UpdateMerchantRequest,
  UpdateMerchantResponse,
} from '../types/merchant'

/** POST /merchants */
export async function createMerchant(
  payload: CreateMerchantRequest,
): Promise<CreateMerchantResponse> {
  const response = await api.post<CreateMerchantResponse>('/merchants', payload)
  return response.data
}

/** GET /merchants */
export async function getMerchants(): Promise<MerchantListResponse> {
  const response = await api.get<MerchantListResponse>('/merchants')
  return response.data
}

/** GET /merchants/:id */
export async function getMerchantById(
  merchantId: number,
): Promise<MerchantResponse> {
  const response = await api.get<MerchantResponse>(`/merchants/${merchantId}`)
  return response.data
}

/** PUT /merchants/:id */
export async function updateMerchant(
  merchantId: number,
  payload: UpdateMerchantRequest,
): Promise<UpdateMerchantResponse> {
  const response = await api.put<UpdateMerchantResponse>(
    `/merchants/${merchantId}`,
    payload,
  )
  return response.data
}

/** PATCH /merchants/:id/commission */
export async function updateCommission(
  merchantId: number,
  payload: UpdateCommissionRequest,
): Promise<UpdateCommissionResponse> {
  const response = await api.patch<UpdateCommissionResponse>(
    `/merchants/${merchantId}/commission`,
    payload,
  )
  return response.data
}

/** DELETE /merchants/:id */
export async function deleteMerchant(
  merchantId: number,
): Promise<DeleteMerchantResponse> {
  const response = await api.delete<DeleteMerchantResponse>(
    `/merchants/${merchantId}`,
  )
  return response.data
}

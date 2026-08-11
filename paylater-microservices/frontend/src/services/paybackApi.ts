import { api } from './api'
import type { PaybackRequest, PaybackResponse } from '../types/payback'

/** POST /payback */
export async function createPayback(
  payload: PaybackRequest,
): Promise<PaybackResponse> {
  const response = await api.post<PaybackResponse>('/payback', payload)
  return response.data
}

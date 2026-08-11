import { api } from './api'
import type {
  CreateTransactionRequest,
  CreateTransactionResponse,
  TransactionListResponse,
  TransactionResponse,
} from '../types/transaction'

/** POST /transactions */
export async function createTransaction(
  payload: CreateTransactionRequest,
): Promise<CreateTransactionResponse> {
  const response = await api.post<CreateTransactionResponse>(
    '/transactions',
    payload,
  )
  return response.data
}

/** GET /transactions */
export async function getTransactions(): Promise<TransactionListResponse> {
  const response = await api.get<TransactionListResponse>('/transactions')
  return response.data
}

/** GET /transactions/:id */
export async function getTransactionById(
  transactionId: number,
): Promise<TransactionResponse> {
  const response = await api.get<TransactionResponse>(
    `/transactions/${transactionId}`,
  )
  return response.data
}

/** GET /transactions/user/:user_id */
export async function getTransactionsByUser(
  userId: number,
): Promise<TransactionListResponse> {
  const response = await api.get<TransactionListResponse>(
    `/transactions/user/${userId}`,
  )
  return response.data
}

/** GET /transactions/merchant/:merchant_id */
export async function getTransactionsByMerchant(
  merchantId: number,
): Promise<TransactionListResponse> {
  const response = await api.get<TransactionListResponse>(
    `/transactions/merchant/${merchantId}`,
  )
  return response.data
}

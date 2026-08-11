import axios from 'axios'

import { getToken } from '../utils/token'

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL

if (!apiBaseUrl) {
  throw new Error(
    'Missing VITE_API_BASE_URL. Set it in frontend/.env (see .env.example).',
  )
}

/**
 * Centralized Axios client for the PayLater API Gateway.
 * Request interceptor attaches Bearer tokens from the token utility.
 */
export const api = axios.create({
  baseURL: apiBaseUrl,
  timeout: 15_000,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
})

api.interceptors.request.use((config) => {
  const token = getToken()

  if (token) {
    config.headers.set('Authorization', `Bearer ${token}`)
  }

  return config
})

export default api

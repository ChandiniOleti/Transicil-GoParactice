import axios from 'axios'
import type { Store } from '@reduxjs/toolkit'

import { logout } from '../features/auth/authSlice'
import { getToken } from '../utils/token'
import { api } from './api'

function isPublicAuthRequest(url: string | undefined, method: string | undefined): boolean {
  if (!url) {
    return false
  }

  const normalizedMethod = method?.toLowerCase()

  if (
    url.endsWith('/login') ||
    url.endsWith('/admin/login') ||
    url.endsWith('/merchant/login')
  ) {
    return true
  }

  return url.endsWith('/users') && normalizedMethod === 'post'
}

let isHandlingUnauthorized = false

/**
 * Registers global Axios response handling (401 → logout + login redirect).
 * Call once after the Redux store is created.
 */
export function setupApiInterceptors(store: Store): void {
  api.interceptors.response.use(
    (response) => response,
    (error: unknown) => {
      if (!axios.isAxiosError(error) || error.response?.status !== 401) {
        return Promise.reject(error)
      }

      const hadToken = getToken() !== null
      const requestUrl = error.config?.url
      const requestMethod = error.config?.method

      if (
        hadToken &&
        !isPublicAuthRequest(requestUrl, requestMethod) &&
        !isHandlingUnauthorized
      ) {
        isHandlingUnauthorized = true
        store.dispatch(logout())

        if (window.location.pathname !== '/login') {
          window.location.assign('/login')
        }

        isHandlingUnauthorized = false
      }

      return Promise.reject(error)
    },
  )
}

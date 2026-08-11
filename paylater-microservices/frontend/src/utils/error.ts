import axios from 'axios'

import type { ApiErrorResponse } from '../types/auth'

function isApiErrorResponse(value: unknown): value is ApiErrorResponse {
  return (
    typeof value === 'object' &&
    value !== null &&
    'error' in value &&
    typeof (value as ApiErrorResponse).error === 'string'
  )
}

/**
 * Extracts a human-readable message from Axios, backend `{ error }`,
 * network, or unknown failures. Does not display UI.
 */
export function getErrorMessage(error: unknown): string {
  if (axios.isAxiosError(error)) {
    const data: unknown = error.response?.data

    if (isApiErrorResponse(data)) {
      return data.error
    }

    if (error.code === 'ECONNABORTED') {
      return 'Request timed out. Please try again.'
    }

    if (!error.response) {
      return 'Network error. Unable to reach the API Gateway.'
    }

    if (typeof error.message === 'string' && error.message.length > 0) {
      return error.message
    }
  }

  if (error instanceof Error && error.message.length > 0) {
    return error.message
  }

  return 'An unexpected error occurred.'
}

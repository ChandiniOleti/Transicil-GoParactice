const ACCESS_TOKEN_KEY = 'paylater_access_token'

export function saveToken(token: string): void {
  if (!token) {
    return
  }
  try {
    localStorage.setItem(ACCESS_TOKEN_KEY, token)
  } catch {
    // Ignore storage write failures (private mode, quota, etc.).
  }
}

export function getToken(): string | null {
  try {
    const token = localStorage.getItem(ACCESS_TOKEN_KEY)
    if (!token || token.trim().length === 0) {
      return null
    }
    return token
  } catch {
    return null
  }
}

export function removeToken(): void {
  try {
    localStorage.removeItem(ACCESS_TOKEN_KEY)
  } catch {
    // Ignore storage remove failures.
  }
}

export function hasToken(): boolean {
  return getToken() !== null
}

export const TOKEN_STORAGE_KEY = ACCESS_TOKEN_KEY

import type { AuthRole, JwtClaims } from '../types/auth'

function base64UrlDecode(value: string): string {
  const normalized = value.replace(/-/g, '+').replace(/_/g, '/')
  const padded = normalized.padEnd(
    normalized.length + ((4 - (normalized.length % 4)) % 4),
    '=',
  )
  return atob(padded)
}

function isAuthRole(value: unknown): value is AuthRole {
  return value === 'USER' || value === 'ADMIN' || value === 'MERCHANT'
}

/** Returns true when the JWT exp claim (Unix seconds) is in the past. */
export function isTokenExpired(exp: number): boolean {
  return exp * 1000 <= Date.now()
}

/**
 * Reads JWT payload claims without verifying the signature.
 * Login APIs return only `{ message, token }`; identity fields live in the token.
 * Signature verification remains for a later authentication step.
 */
export function readJwtClaims(token: string): JwtClaims | null {
  try {
    const parts = token.split('.')
    if (parts.length < 2 || !parts[1]) {
      return null
    }

    const parsed: unknown = JSON.parse(base64UrlDecode(parts[1]))
    if (typeof parsed !== 'object' || parsed === null) {
      return null
    }

    const record = parsed as Record<string, unknown>
    const userId = record.user_id
    const email = record.email
    const role = record.role
    const exp = record.exp

    if (
      typeof userId !== 'number' ||
      typeof email !== 'string' ||
      !isAuthRole(role) ||
      typeof exp !== 'number'
    ) {
      return null
    }

    return {
      user_id: userId,
      email,
      role,
      exp,
    }
  } catch {
    return null
  }
}

import type { AuthRole } from '../types/auth'

export function getDashboardPath(role: AuthRole): string {
  switch (role) {
    case 'ADMIN':
      return '/admin/dashboard'
    case 'MERCHANT':
      return '/merchant/dashboard'
    case 'USER':
      return '/user/dashboard'
  }
}

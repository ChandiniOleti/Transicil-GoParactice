import { Navigate, Outlet } from 'react-router-dom'

import { useAppSelector } from '../app/hooks'
import {
  selectAuthRole,
  selectIsAuthenticated,
} from '../features/auth/authSelectors'
import type { AuthRole } from '../types/auth'
import { getDashboardPath } from '../utils/routes'

export interface RoleRouteProps {
  allowedRoles: readonly AuthRole[]
}

/**
 * Allows nested routes only when the authenticated role is permitted.
 * Wrong role redirects to that role's dashboard.
 */
export default function RoleRoute({ allowedRoles }: RoleRouteProps) {
  const isAuthenticated = useAppSelector(selectIsAuthenticated)
  const role = useAppSelector(selectAuthRole)

  if (!isAuthenticated || role === null) {
    return <Navigate to="/login" replace />
  }

  if (!allowedRoles.includes(role)) {
    return <Navigate to={getDashboardPath(role)} replace />
  }

  return <Outlet />
}

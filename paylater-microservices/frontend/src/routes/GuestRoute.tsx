import { Navigate, Outlet } from 'react-router-dom'

import { useAppSelector } from '../app/hooks'
import {
  selectAuthRole,
  selectAuthStatus,
  selectIsAuthenticated,
} from '../features/auth/authSelectors'
import { getDashboardPath } from '../utils/routes'

/** Public routes; authenticated users are sent to their dashboard. */
export default function GuestRoute() {
  const isAuthenticated = useAppSelector(selectIsAuthenticated)
  const status = useAppSelector(selectAuthStatus)
  const role = useAppSelector(selectAuthRole)

  if (isAuthenticated && status === 'authenticated' && role !== null) {
    return <Navigate to={getDashboardPath(role)} replace />
  }

  return <Outlet />
}

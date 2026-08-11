import { Navigate, Outlet, useLocation } from 'react-router-dom'

import { useAppSelector } from '../app/hooks'
import {
  selectAuthStatus,
  selectIsAuthenticated,
} from '../features/auth/authSelectors'

/**
 * Requires Redux-authenticated identity.
 * Token-only / idle state is not treated as authenticated.
 */
export default function ProtectedRoute() {
  const isAuthenticated = useAppSelector(selectIsAuthenticated)
  const status = useAppSelector(selectAuthStatus)
  const location = useLocation()

  if (!isAuthenticated || status !== 'authenticated') {
    return <Navigate to="/login" replace state={{ from: location }} />
  }

  return <Outlet />
}

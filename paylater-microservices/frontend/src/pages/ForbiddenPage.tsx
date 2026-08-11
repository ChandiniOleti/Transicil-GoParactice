import { Link } from 'react-router-dom'

import { useAppSelector } from '../app/hooks'
import { selectAuthRole, selectIsAuthenticated } from '../features/auth/authSelectors'
import { getDashboardPath } from '../utils/routes'

export default function ForbiddenPage() {
  const isAuthenticated = useAppSelector(selectIsAuthenticated)
  const role = useAppSelector(selectAuthRole)

  const homePath =
    isAuthenticated && role !== null ? getDashboardPath(role) : '/login'

  return (
    <div>
      <h1>403</h1>
      <h2>Access denied</h2>
      <p>You do not have permission to access this page.</p>
      <Link to={homePath}>
        {isAuthenticated ? 'Back to Dashboard' : 'Back to Login'}
      </Link>
    </div>
  )
}

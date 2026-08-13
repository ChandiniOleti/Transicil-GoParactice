import { useCallback, useEffect, useState } from 'react'
import { Link } from 'react-router-dom'

import { useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import FinancialSummary from '../../features/users/FinancialSummary'
import UserSummary from '../../features/users/UserSummary'
import { getUserById } from '../../services/userApi'
import type { User } from '../../types/user'
import { getErrorMessage } from '../../utils/error'

export default function UserDashboardPage() {
  const authUser = useAppSelector(selectCurrentUser)
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadUser = useCallback(async () => {
    if (!authUser) {
      setUser(null)
      setError('Authenticated user information is unavailable.')
      return
    }

    setIsLoading(true)
    setError(null)

    try {
      const response = await getUserById(authUser.user_id)
      setUser(response)
    } catch (err) {
      setUser(null)
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [authUser])

  useEffect(() => {
    void loadUser()
  }, [loadUser])

  return (
    <PageContainer
      title="User Dashboard"
      description="Your PayLater account overview."
    >
      {isLoading ? <Loader label="Loading account details" /> : null}

      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load dashboard" message={error} />
          <Button type="button" variant="secondary" onClick={() => void loadUser()}>
            Retry
          </Button>
        </div>
      ) : null}

      {!isLoading && !error && user ? (
        <div className="pl-user-dashboard">
          <UserSummary user={user} />
          <FinancialSummary user={user} />
          <div className="pl-user-profile__actions">
            <Link
              to="/user/purchase"
              className="pl-button pl-button--primary pl-button--medium"
            >
              Make Purchase
            </Link>
            <Link
              to="/user/transactions"
              className="pl-button pl-button--secondary pl-button--medium"
            >
              View Transactions
            </Link>
          </div>
        </div>
      ) : null}
    </PageContainer>
  )
}

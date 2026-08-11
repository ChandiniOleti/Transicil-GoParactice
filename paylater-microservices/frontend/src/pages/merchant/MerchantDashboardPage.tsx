import { useCallback, useEffect, useState } from 'react'

import { useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import MerchantStats from '../../features/merchants/MerchantStats'
import MerchantSummary from '../../features/merchants/MerchantSummary'
import { getMerchantById } from '../../services/merchantApi'
import type { Merchant } from '../../types/merchant'
import { getErrorMessage } from '../../utils/error'

export default function MerchantDashboardPage() {
  const authUser = useAppSelector(selectCurrentUser)
  const [merchant, setMerchant] = useState<Merchant | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadMerchant = useCallback(async () => {
    if (!authUser) {
      setMerchant(null)
      setError('Authenticated merchant information is unavailable.')
      return
    }

    setIsLoading(true)
    setError(null)

    try {
      const response = await getMerchantById(authUser.user_id)
      setMerchant(response)
    } catch (err) {
      setMerchant(null)
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [authUser])

  useEffect(() => {
    void loadMerchant()
  }, [loadMerchant])

  return (
    <PageContainer
      title="Merchant Dashboard"
      description="Your PayLater merchant account overview."
    >
      {isLoading ? <Loader label="Loading merchant details" /> : null}

      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load dashboard" message={error} />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadMerchant()}
          >
            Retry
          </Button>
        </div>
      ) : null}

      {!isLoading && !error && merchant ? (
        <div className="pl-merchant-dashboard">
          <MerchantSummary merchant={merchant} />
          <MerchantStats merchant={merchant} />
        </div>
      ) : null}
    </PageContainer>
  )
}

import { useCallback, useEffect, useState } from 'react'

import { useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import ProfileField from '../../components/ui/ProfileField'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import { getMerchantById } from '../../services/merchantApi'
import type { Merchant } from '../../types/merchant'
import { formatCommission } from '../../utils/commission'
import { getErrorMessage } from '../../utils/error'

export default function MerchantProfilePage() {
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
      title="Merchant Profile"
      description="Read-only view of your PayLater merchant account."
    >
      {isLoading ? <Loader label="Loading profile" /> : null}

      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load profile" message={error} />
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
        <div className="pl-user-profile">
          <header className="pl-user-profile__header">
            <h2 className="pl-user-profile__name">{merchant.merchant_name}</h2>
            <p className="pl-user-profile__email">{merchant.email}</p>
          </header>

          <div className="pl-user-profile__sections">
            <Card title="Business Information">
              <dl className="pl-profile-fields">
                <ProfileField
                  label="Merchant Name"
                  value={merchant.merchant_name}
                />
                <ProfileField label="Email" value={merchant.email} />
                <ProfileField label="Phone" value={merchant.phone} />
                <ProfileField
                  label="Merchant ID"
                  value={String(merchant.id)}
                />
              </dl>
            </Card>

            <Card title="Commission Information">
              <dl className="pl-profile-fields">
                <ProfileField
                  label="Commission Rate"
                  value={formatCommission(merchant.commission)}
                />
              </dl>
            </Card>
          </div>
        </div>
      ) : null}
    </PageContainer>
  )
}

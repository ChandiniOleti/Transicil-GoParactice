import { useCallback, useEffect, useState } from 'react'

import { useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import MerchantCustomerDetails from '../../features/merchants/MerchantCustomerDetails'
import MerchantRecentTransactions from '../../features/merchants/MerchantRecentTransactions'
import MerchantStats from '../../features/merchants/MerchantStats'
import MerchantSummary from '../../features/merchants/MerchantSummary'
import {
  buildCustomerDetails,
  getRecentTransactions,
} from '../../features/merchants/merchantDashboardUtils'
import { getMerchantById } from '../../services/merchantApi'
import { getTransactionsByMerchant } from '../../services/transactionApi'
import type { Merchant } from '../../types/merchant'
import type { Transaction } from '../../types/transaction'
import { getErrorMessage } from '../../utils/error'

export default function MerchantDashboardPage() {
  const authUser = useAppSelector(selectCurrentUser)
  const [merchant, setMerchant] = useState<Merchant | null>(null)
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [merchantError, setMerchantError] = useState<string | null>(null)
  const [transactionsError, setTransactionsError] = useState<string | null>(null)

  const loadDashboard = useCallback(async () => {
    if (!authUser) {
      setMerchant(null)
      setTransactions([])
      setMerchantError('Authenticated merchant information is unavailable.')
      setTransactionsError(null)
      return
    }

    setIsLoading(true)
    setMerchantError(null)
    setTransactionsError(null)

    const [merchantResult, transactionsResult] = await Promise.allSettled([
      getMerchantById(authUser.user_id),
      getTransactionsByMerchant(authUser.user_id),
    ])

    if (merchantResult.status === 'fulfilled') {
      setMerchant(merchantResult.value)
    } else {
      setMerchant(null)
      setMerchantError(getErrorMessage(merchantResult.reason))
    }

    if (transactionsResult.status === 'fulfilled') {
      setTransactions(transactionsResult.value)
    } else {
      setTransactions([])
      setTransactionsError(getErrorMessage(transactionsResult.reason))
    }

    setIsLoading(false)
  }, [authUser])

  useEffect(() => {
    void loadDashboard()
  }, [loadDashboard])

  const recentTransactions = getRecentTransactions(transactions)
  const customerDetails = merchant
    ? buildCustomerDetails(transactions, merchant.commission)
    : []

  return (
    <PageContainer
      title="Merchant Dashboard"
      description="Your PayLater merchant account overview."
    >
      {isLoading ? <Loader label="Loading merchant dashboard" /> : null}

      {!isLoading && merchantError ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load dashboard" message={merchantError} />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadDashboard()}
          >
            Retry
          </Button>
        </div>
      ) : null}

      {!isLoading && !merchantError && merchant ? (
        <div className="pl-merchant-dashboard">
          <MerchantSummary merchant={merchant} />

          {!transactionsError ? (
            <>
              <MerchantStats merchant={merchant} transactions={transactions} />
              <MerchantCustomerDetails customers={customerDetails} />
              <MerchantRecentTransactions transactions={recentTransactions} />
            </>
          ) : (
            <div className="pl-dashboard-error">
              <ErrorMessage
                title="Unable to load transactions"
                message={transactionsError}
              />
              <Button
                type="button"
                variant="secondary"
                onClick={() => void loadDashboard()}
              >
                Retry
              </Button>
            </div>
          )}
        </div>
      ) : null}
    </PageContainer>
  )
}

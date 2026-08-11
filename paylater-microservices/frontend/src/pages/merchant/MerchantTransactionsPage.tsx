import { useCallback, useEffect, useState } from 'react'

import { useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import PageContainer from '../../components/layout/PageContainer'
import Table, { type TableColumn } from '../../components/ui/Table'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import { getTransactionsByMerchant } from '../../services/transactionApi'
import type { Transaction } from '../../types/transaction'
import { formatCurrency } from '../../utils/currency'
import { formatDateTime } from '../../utils/date'
import { getErrorMessage } from '../../utils/error'

const transactionColumns: TableColumn<Transaction>[] = [
  { key: 'id', header: 'Transaction ID' },
  {
    key: 'user_id',
    header: 'Customer ID',
    render: (row) => String(row.user_id),
  },
  {
    key: 'amount',
    header: 'Amount',
    render: (row) => formatCurrency(row.amount),
  },
  { key: 'transaction_type', header: 'Type' },
  {
    key: 'transaction_date',
    header: 'Date',
    render: (row) =>
      row.transaction_date ? formatDateTime(row.transaction_date) : '—',
  },
]

export default function MerchantTransactionsPage() {
  const authUser = useAppSelector(selectCurrentUser)
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadTransactions = useCallback(async () => {
    if (!authUser) {
      setTransactions([])
      setError('Authenticated merchant information is unavailable.')
      return
    }

    setIsLoading(true)
    setError(null)

    try {
      const response = await getTransactionsByMerchant(authUser.user_id)
      setTransactions(response)
    } catch (err) {
      setTransactions([])
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [authUser])

  useEffect(() => {
    void loadTransactions()
  }, [loadTransactions])

  return (
    <PageContainer
      title="Transactions"
      description="PayLater transactions for your merchant account."
    >
      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage
            title="Unable to load transactions"
            message={error}
          />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadTransactions()}
          >
            Retry
          </Button>
        </div>
      ) : (
        <Table
          columns={transactionColumns}
          data={transactions}
          loading={isLoading}
          emptyMessage="No transactions found."
        />
      )}
    </PageContainer>
  )
}

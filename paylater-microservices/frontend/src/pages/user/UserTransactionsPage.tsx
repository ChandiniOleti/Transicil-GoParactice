import { useCallback, useEffect, useState } from 'react'

import { useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import PageContainer from '../../components/layout/PageContainer'
import Table, { type TableColumn } from '../../components/ui/Table'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import { getTransactionsByUser } from '../../services/transactionApi'
import type { Transaction } from '../../types/transaction'
import { formatCurrency } from '../../utils/currency'
import { formatDateTime } from '../../utils/date'
import { getErrorMessage } from '../../utils/error'

const transactionColumns: TableColumn<Transaction>[] = [
  { key: 'id', header: 'Transaction ID' },
  {
    key: 'amount',
    header: 'Amount',
    render: (row) => formatCurrency(row.amount),
  },
  { key: 'transaction_type', header: 'Type' },
  {
    key: 'merchant_id',
    header: 'Merchant',
    render: (row) =>
      row.merchant_id !== undefined ? `Merchant #${row.merchant_id}` : '—',
  },
  {
    key: 'transaction_date',
    header: 'Date',
    render: (row) =>
      row.transaction_date ? formatDateTime(row.transaction_date) : '—',
  },
]

export default function UserTransactionsPage() {
  const authUser = useAppSelector(selectCurrentUser)
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [isLoadingTransactions, setIsLoadingTransactions] = useState(false)
  const [transactionsError, setTransactionsError] = useState<string | null>(null)

  const loadTransactions = useCallback(async () => {
    if (!authUser) {
      setTransactions([])
      setTransactionsError('Authenticated user information is unavailable.')
      return
    }

    setIsLoadingTransactions(true)
    setTransactionsError(null)

    try {
      const response = await getTransactionsByUser(authUser.user_id)
      setTransactions(response)
    } catch (err) {
      setTransactions([])
      setTransactionsError(getErrorMessage(err))
    } finally {
      setIsLoadingTransactions(false)
    }
  }, [authUser])

  useEffect(() => {
    void loadTransactions()
  }, [loadTransactions])

  return (
    <PageContainer
      title="Transactions"
      description="Your PayLater transaction history."
    >
      <div className="pl-user-transactions">
        {!isLoadingTransactions && transactionsError ? (
          <div className="pl-dashboard-error">
            <ErrorMessage
              title="Unable to load transactions"
              message={transactionsError}
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
            loading={isLoadingTransactions}
            emptyMessage="No transactions found."
          />
        )}
      </div>
    </PageContainer>
  )
}

import { useCallback, useEffect, useState } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import StatCard from '../../components/ui/StatCard'
import Table, { type TableColumn } from '../../components/ui/Table'
import {
  countPaybacks,
  countPurchases,
  sumFeesCollected,
} from '../../features/admin/adminTransactionUtils'
import { getTransactions } from '../../services/transactionApi'
import type { Transaction } from '../../types/transaction'
import { formatCurrency } from '../../utils/currency'
import { formatDateTime } from '../../utils/date'
import { getErrorMessage } from '../../utils/error'

const transactionColumns: TableColumn<Transaction>[] = [
  { key: 'id', header: 'Transaction ID' },
  {
    key: 'user_id',
    header: 'User ID',
    render: (row) => String(row.user_id),
  },
  {
    key: 'merchant_id',
    header: 'Merchant ID',
    render: (row) =>
      row.merchant_id !== undefined ? String(row.merchant_id) : '—',
  },
  {
    key: 'amount',
    header: 'Amount',
    render: (row) => formatCurrency(row.amount),
  },
  { key: 'transaction_type', header: 'Transaction Type' },
  {
    key: 'transaction_date',
    header: 'Transaction Date',
    render: (row) =>
      row.transaction_date ? formatDateTime(row.transaction_date) : '—',
  },
]

export default function AdminTransactionsPage() {
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadTransactions = useCallback(async () => {
    setIsLoading(true)
    setError(null)

    try {
      const response = await getTransactions()
      setTransactions(response)
    } catch (err) {
      setTransactions([])
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [])

  useEffect(() => {
    void loadTransactions()
  }, [loadTransactions])

  return (
    <PageContainer
      title="Transactions"
      description="View all PayLater platform transactions."
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
        <>
          {!isLoading ? (
            <section
              className="pl-admin-stats pl-admin-transactions-stats"
              aria-label="Transaction summary statistics"
            >
              <StatCard
                label="Total Purchases"
                value={String(countPurchases(transactions))}
              />
              <StatCard
                label="Total Paybacks"
                value={String(countPaybacks(transactions))}
              />
              <StatCard
                label="Fees Collected"
                value={formatCurrency(sumFeesCollected(transactions).toFixed(2))}
              />
            </section>
          ) : null}

          <Card title="All Transactions">
          <Table
            columns={transactionColumns}
            data={transactions}
            loading={isLoading}
            emptyMessage="No transactions found."
          />
        </Card>
        </>
      )}
    </PageContainer>
  )
}

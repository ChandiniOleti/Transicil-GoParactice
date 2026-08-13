import { Link } from 'react-router-dom'

import Card from '../../components/ui/Card'
import EmptyState from '../../components/common/EmptyState'
import Table, { type TableColumn } from '../../components/ui/Table'
import type { Transaction } from '../../types/transaction'
import { formatCurrency } from '../../utils/currency'
import { formatDateTime } from '../../utils/date'
import { RECENT_TRANSACTION_LIMIT } from './merchantDashboardUtils'

export interface MerchantRecentTransactionsProps {
  transactions: Transaction[]
}

const recentTransactionColumns: TableColumn<Transaction>[] = [
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

export default function MerchantRecentTransactions({
  transactions,
}: MerchantRecentTransactionsProps) {
  return (
    <Card title="Recent Transactions" className="pl-merchant-recent-transactions">
      <div className="pl-merchant-recent-transactions__header">
        <p className="pl-merchant-recent-transactions__subtitle">
          Latest {RECENT_TRANSACTION_LIMIT} transactions for your merchant
          account.
        </p>
        <Link
          to="/merchant/transactions"
          className="pl-button pl-button--outline pl-button--small"
        >
          View All Transactions
        </Link>
      </div>

      {transactions.length === 0 ? (
        <EmptyState title="No transactions found." />
      ) : (
        <Table
          columns={recentTransactionColumns}
          data={transactions}
          emptyMessage="No transactions found."
        />
      )}
    </Card>
  )
}

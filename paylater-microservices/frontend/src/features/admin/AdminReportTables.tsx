import Card from '../../components/ui/Card'
import Table, { type TableColumn } from '../../components/ui/Table'
import type { UserReport } from '../../types/report'
import { formatCommission } from '../../utils/commission'
import { formatCurrency } from '../../utils/currency'

export interface MerchantFeeRow {
  merchantId: number
  merchantName: string
  commission: string
  totalFeeCollected: string
}

export interface AdminReportTablesProps {
  usersWithDue: UserReport[]
  creditLimitUsers: UserReport[]
  merchantFees: MerchantFeeRow[]
}

const userReportColumns: TableColumn<UserReport>[] = [
  { key: 'id', header: 'User ID' },
  { key: 'name', header: 'Name' },
  { key: 'email', header: 'Email' },
  {
    key: 'current_due',
    header: 'Current Due',
    render: (row) => formatCurrency(row.current_due),
  },
  {
    key: 'credit_limit',
    header: 'Credit Limit',
    render: (row) => formatCurrency(row.credit_limit),
  },
]

const merchantFeeColumns: TableColumn<MerchantFeeRow>[] = [
  { key: 'merchantId', header: 'Merchant ID' },
  { key: 'merchantName', header: 'Merchant Name' },
  {
    key: 'commission',
    header: 'Commission',
    render: (row) => formatCommission(row.commission),
  },
  {
    key: 'totalFeeCollected',
    header: 'Total Fee Collected',
    render: (row) => formatCurrency(row.totalFeeCollected),
  },
]

export default function AdminReportTables({
  usersWithDue,
  creditLimitUsers,
  merchantFees,
}: AdminReportTablesProps) {
  return (
    <div className="pl-admin-report-tables">
      <Card title="Users with Outstanding Dues">
        <Table
          columns={userReportColumns}
          data={usersWithDue}
          emptyMessage="No users with outstanding dues."
        />
      </Card>

      <Card title="Users at Credit Limit">
        <Table
          columns={userReportColumns}
          data={creditLimitUsers}
          emptyMessage="No users at credit limit."
        />
      </Card>

      <Card title="Merchant Commission Summary">
        <Table
          columns={merchantFeeColumns}
          data={merchantFees}
          emptyMessage="No merchant commission data available."
        />
      </Card>
    </div>
  )
}

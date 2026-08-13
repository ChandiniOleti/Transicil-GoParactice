import Card from '../../components/ui/Card'
import Table, { type TableColumn } from '../../components/ui/Table'
import { formatCurrency } from '../../utils/currency'
import type { MerchantCustomerRow } from './merchantDashboardUtils'

export interface MerchantCustomerDetailsProps {
  customers: MerchantCustomerRow[]
}

const customerColumns: TableColumn<MerchantCustomerRow>[] = [
  { key: 'customerName', header: 'Customer Name' },
  {
    key: 'purchaseCount',
    header: 'Purchase Count',
    render: (row) => String(row.purchaseCount),
  },
  {
    key: 'totalPurchasedAmount',
    header: 'Total Purchased Amount',
    render: (row) => formatCurrency(row.totalPurchasedAmount.toFixed(2)),
  },
  {
    key: 'commission',
    header: 'Commission',
    render: (row) => formatCurrency(row.commission.toFixed(2)),
  },
]

export default function MerchantCustomerDetails({
  customers,
}: MerchantCustomerDetailsProps) {
  return (
    <Card title="Customer Details">
      <Table
        columns={customerColumns}
        data={customers}
        emptyMessage="No customer purchases yet."
      />
    </Card>
  )
}

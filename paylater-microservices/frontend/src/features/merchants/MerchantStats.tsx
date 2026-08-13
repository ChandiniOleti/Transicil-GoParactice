import StatCard from '../../components/ui/StatCard'
import { formatCurrency } from '../../utils/currency'
import {
  calculateNetAmount,
  calculateTotalCommission,
  calculateTotalCustomers,
  calculateTotalPurchases,
  calculateTotalSales,
} from './merchantDashboardUtils'
import type { Transaction } from '../../types/transaction'
import type { Merchant } from '../../types/merchant'

export interface MerchantStatsProps {
  merchant: Merchant
  transactions: Transaction[]
}

export default function MerchantStats({
  merchant,
  transactions,
}: MerchantStatsProps) {
  const totalSales = calculateTotalSales(transactions)
  const totalCommission = calculateTotalCommission(
    transactions,
    merchant.commission,
  )
  const netAmount = calculateNetAmount(totalSales, totalCommission)

  return (
    <section className="pl-merchant-stats" aria-label="Merchant statistics">
      <StatCard
        label="Total Customers"
        value={String(calculateTotalCustomers(transactions))}
      />
      <StatCard
        label="Total Purchases"
        value={String(calculateTotalPurchases(transactions))}
      />
      <StatCard
        label="Total Commission"
        value={formatCurrency(totalCommission.toFixed(2))}
      />
      <StatCard label="Net Amount" value={formatCurrency(netAmount.toFixed(2))} />
    </section>
  )
}

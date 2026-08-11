import StatCard from '../../components/ui/StatCard'
import { formatCurrency } from '../../utils/currency'

export interface AdminSummaryStatsProps {
  totalDues: string
  usersWithDueCount: number
  creditLimitUsersCount: number
  merchantCount: number
  totalCommissionCollected: string
}

export default function AdminSummaryStats({
  totalDues,
  usersWithDueCount,
  creditLimitUsersCount,
  merchantCount,
  totalCommissionCollected,
}: AdminSummaryStatsProps) {
  return (
    <section className="pl-admin-stats" aria-label="Admin summary statistics">
      <StatCard
        label="Total Outstanding Dues"
        value={formatCurrency(totalDues)}
      />
      <StatCard
        label="Users with Due"
        value={String(usersWithDueCount)}
      />
      <StatCard
        label="Users at Credit Limit"
        value={String(creditLimitUsersCount)}
      />
      <StatCard label="Total Merchants" value={String(merchantCount)} />
      <StatCard
        label="Total Commission Collected"
        value={formatCurrency(totalCommissionCollected)}
      />
    </section>
  )
}

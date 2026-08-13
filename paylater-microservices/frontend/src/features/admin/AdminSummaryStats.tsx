import StatCard from '../../components/ui/StatCard'
import { formatCurrency } from '../../utils/currency'

export type AdminSummaryStatKey =
  | 'totalDues'
  | 'usersWithDue'
  | 'creditLimitUsers'
  | 'merchantCount'
  | 'totalCommissionCollected'

export interface AdminSummaryStatsProps {
  totalDues: string
  usersWithDueCount: number
  creditLimitUsersCount: number
  merchantCount: number
  totalCommissionCollected: string
  clickableStats?: Partial<Record<AdminSummaryStatKey, boolean>>
  onStatClick?: (key: AdminSummaryStatKey) => void
}

export default function AdminSummaryStats({
  totalDues,
  usersWithDueCount,
  creditLimitUsersCount,
  merchantCount,
  totalCommissionCollected,
  clickableStats,
  onStatClick,
}: AdminSummaryStatsProps) {
  function statProps(key: AdminSummaryStatKey) {
    const clickable = Boolean(clickableStats?.[key] && onStatClick)

    return {
      clickable,
      onClick: clickable ? () => onStatClick?.(key) : undefined,
    }
  }

  return (
    <section className="pl-admin-stats" aria-label="Admin summary statistics">
      <StatCard
        label="Total Outstanding Dues"
        value={formatCurrency(totalDues)}
        {...statProps('totalDues')}
      />
      <StatCard
        label="Users with Due"
        value={String(usersWithDueCount)}
        {...statProps('usersWithDue')}
      />
      <StatCard
        label="Users at Credit Limit"
        value={String(creditLimitUsersCount)}
        {...statProps('creditLimitUsers')}
      />
      <StatCard
        label="Total Merchants"
        value={String(merchantCount)}
        {...statProps('merchantCount')}
      />
      <StatCard
        label="Total Commission Collected"
        value={formatCurrency(totalCommissionCollected)}
        {...statProps('totalCommissionCollected')}
      />
    </section>
  )
}

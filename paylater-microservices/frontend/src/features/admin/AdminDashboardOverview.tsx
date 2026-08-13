import StatCard from '../../components/ui/StatCard'
import { formatCurrency } from '../../utils/currency'

export interface AdminDashboardOverviewProps {
  totalUsers: number
  totalMerchants: number
  totalOutstandingDue: string
}

export default function AdminDashboardOverview({
  totalUsers,
  totalMerchants,
  totalOutstandingDue,
}: AdminDashboardOverviewProps) {
  return (
    <section
      className="pl-admin-stats pl-admin-dashboard-overview"
      aria-label="Admin dashboard overview"
    >
      <StatCard label="Total Users" value={String(totalUsers)} />
      <StatCard label="Total Merchants" value={String(totalMerchants)} />
      <StatCard
        label="Total Outstanding Due"
        value={formatCurrency(totalOutstandingDue)}
      />
    </section>
  )
}

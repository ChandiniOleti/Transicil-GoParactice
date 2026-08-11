import StatCard from '../../components/ui/StatCard'
import type { User } from '../../types/user'
import {
  calculateAvailableCredit,
  formatCurrency,
} from '../../utils/currency'

export interface FinancialSummaryProps {
  user: User
}

export default function FinancialSummary({ user }: FinancialSummaryProps) {
  const availableCredit = calculateAvailableCredit(
    user.credit_limit,
    user.current_due,
  )

  return (
    <section className="pl-financial-summary" aria-label="Financial summary">
      <StatCard label="Credit Limit" value={formatCurrency(user.credit_limit)} />
      <StatCard label="Current Due" value={formatCurrency(user.current_due)} />
      <StatCard
        label="Available Credit"
        value={formatCurrency(availableCredit)}
      />
    </section>
  )
}

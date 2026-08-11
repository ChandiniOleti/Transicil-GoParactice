import StatCard from '../../components/ui/StatCard'
import type { Merchant } from '../../types/merchant'
import { formatCommission } from '../../utils/commission'

export interface MerchantStatsProps {
  merchant: Merchant
}

export default function MerchantStats({ merchant }: MerchantStatsProps) {
  return (
    <section className="pl-merchant-stats" aria-label="Merchant statistics">
      <StatCard
        label="Commission Rate"
        value={formatCommission(merchant.commission)}
      />
    </section>
  )
}

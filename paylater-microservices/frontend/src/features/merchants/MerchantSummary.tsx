import Card from '../../components/ui/Card'
import ProfileField from '../../components/ui/ProfileField'
import type { Merchant } from '../../types/merchant'

export interface MerchantSummaryProps {
  merchant: Merchant
}

export default function MerchantSummary({ merchant }: MerchantSummaryProps) {
  return (
    <Card title="Merchant Account" className="pl-merchant-summary">
      <dl className="pl-profile-fields">
        <ProfileField label="Merchant Name" value={merchant.merchant_name} />
        <ProfileField label="Email" value={merchant.email} />
        <ProfileField label="Phone" value={merchant.phone} />
        <ProfileField label="Merchant ID" value={String(merchant.id)} />
      </dl>
    </Card>
  )
}

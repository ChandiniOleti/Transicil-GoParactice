import type { MerchantFeeRow } from './AdminReportTables'
import { getMerchantFeeReport, getUserDueReport } from '../../services/reportApi'
import type { Merchant } from '../../types/merchant'
import type { UserReport } from '../../types/report'

export function sumCommissionCollected(
  fees: Array<{ total_fee_collected: string }>,
): string {
  const total = fees.reduce((sum, fee) => {
    const value = Number(fee.total_fee_collected)
    return sum + (Number.isFinite(value) ? value : 0)
  }, 0)

  return total.toFixed(2)
}

export async function buildMerchantFeeRows(
  merchants: Merchant[],
): Promise<MerchantFeeRow[]> {
  const feeReports = await Promise.all(
    merchants.map((merchant) => getMerchantFeeReport(merchant.id)),
  )

  return merchants.map((merchant, index) => ({
    merchantId: merchant.id,
    merchantName: merchant.merchant_name,
    commission: merchant.commission,
    totalFeeCollected: feeReports[index]?.total_fee_collected ?? '0.00',
  }))
}

export async function buildUserDueDetails(
  userIds: number[],
): Promise<UserReport[]> {
  if (userIds.length === 0) {
    return []
  }

  return Promise.all(userIds.map((userId) => getUserDueReport(userId)))
}

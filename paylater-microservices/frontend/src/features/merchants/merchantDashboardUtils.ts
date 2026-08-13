import type { Transaction } from '../../types/transaction'

function parseAmount(value: string): number {
  const parsed = Number(value)
  return Number.isFinite(parsed) ? parsed : 0
}

function purchaseTransactions(transactions: Transaction[]): Transaction[] {
  return transactions.filter(
    (transaction) => transaction.transaction_type === 'PURCHASE',
  )
}

export function calculateTotalCustomers(transactions: Transaction[]): number {
  const userIds = new Set(
    purchaseTransactions(transactions).map((transaction) => transaction.user_id),
  )
  return userIds.size
}

export function calculateTotalPurchases(transactions: Transaction[]): number {
  return purchaseTransactions(transactions).length
}

export function calculateTotalSales(transactions: Transaction[]): number {
  return purchaseTransactions(transactions).reduce(
    (sum, transaction) => sum + parseAmount(transaction.amount),
    0,
  )
}

export function calculateTotalCommission(
  transactions: Transaction[],
  commissionRate: string,
): number {
  const rate = Number(commissionRate)

  return purchaseTransactions(transactions).reduce((sum, transaction) => {
    const commissionAmount = parseAmount(transaction.commission_amount)
    if (commissionAmount > 0) {
      return sum + commissionAmount
    }

    if (!Number.isFinite(rate)) {
      return sum
    }

    const amount = parseAmount(transaction.amount)
    return sum + (amount * rate) / 100
  }, 0)
}

export function calculateNetAmount(
  totalSales: number,
  totalCommission: number,
): number {
  return totalSales - totalCommission
}

export interface MerchantCustomerRow {
  userId: number
  customerName: string
  purchaseCount: number
  totalPurchasedAmount: number
  commission: number
}

export function buildCustomerDetails(
  transactions: Transaction[],
  commissionRate: string,
): MerchantCustomerRow[] {
  const rate = Number(commissionRate)
  const grouped = new Map<
    number,
    { purchaseCount: number; totalPurchasedAmount: number; commission: number }
  >()

  for (const transaction of purchaseTransactions(transactions)) {
    const existing = grouped.get(transaction.user_id) ?? {
      purchaseCount: 0,
      totalPurchasedAmount: 0,
      commission: 0,
    }
    const amount = parseAmount(transaction.amount)
    let commission = parseAmount(transaction.commission_amount)

    if (commission <= 0 && Number.isFinite(rate)) {
      commission = (amount * rate) / 100
    }

    grouped.set(transaction.user_id, {
      purchaseCount: existing.purchaseCount + 1,
      totalPurchasedAmount: existing.totalPurchasedAmount + amount,
      commission: existing.commission + commission,
    })
  }

  return Array.from(grouped.entries())
    .map(([userId, row]) => ({
      userId,
      customerName: `Customer #${userId}`,
      purchaseCount: row.purchaseCount,
      totalPurchasedAmount: row.totalPurchasedAmount,
      commission: row.commission,
    }))
    .sort((left, right) => left.customerName.localeCompare(right.customerName))
}

const RECENT_TRANSACTION_LIMIT = 5

export function getRecentTransactions(
  transactions: Transaction[],
  limit = RECENT_TRANSACTION_LIMIT,
): Transaction[] {
  return [...transactions]
    .sort((left, right) => {
      const leftTime = left.transaction_date
        ? Date.parse(left.transaction_date)
        : Number.NaN
      const rightTime = right.transaction_date
        ? Date.parse(right.transaction_date)
        : Number.NaN

      if (Number.isFinite(leftTime) && Number.isFinite(rightTime)) {
        return rightTime - leftTime
      }

      if (Number.isFinite(leftTime)) {
        return -1
      }

      if (Number.isFinite(rightTime)) {
        return 1
      }

      return right.id - left.id
    })
    .slice(0, limit)
}

export { RECENT_TRANSACTION_LIMIT }

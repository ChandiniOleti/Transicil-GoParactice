import type { Transaction } from '../../types/transaction'

function parseAmount(value: string): number {
  const parsed = Number(value)
  return Number.isFinite(parsed) ? parsed : 0
}

export function countPurchases(transactions: Transaction[]): number {
  return transactions.filter(
    (transaction) => transaction.transaction_type === 'PURCHASE',
  ).length
}

export function countPaybacks(transactions: Transaction[]): number {
  return transactions.filter(
    (transaction) => transaction.transaction_type === 'PAYBACK',
  ).length
}

export function sumFeesCollected(transactions: Transaction[]): number {
  return transactions
    .filter((transaction) => transaction.transaction_type === 'PURCHASE')
    .reduce(
      (sum, transaction) => sum + parseAmount(transaction.commission_amount),
      0,
    )
}

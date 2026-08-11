const inrFormatter = new Intl.NumberFormat('en-IN', {
  style: 'currency',
  currency: 'INR',
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
})

/**
 * Formats DECIMAL string/number values as INR for display.
 */
export function formatCurrency(value: string | number): string {
  const amount = typeof value === 'number' ? value : Number(value)

  if (!Number.isFinite(amount)) {
    return inrFormatter.format(0)
  }

  return inrFormatter.format(amount)
}

/**
 * Computes available credit from DECIMAL string fields.
 */
export function calculateAvailableCredit(
  creditLimit: string,
  currentDue: string,
): number {
  const limit = Number(creditLimit)
  const due = Number(currentDue)

  if (!Number.isFinite(limit) || !Number.isFinite(due)) {
    return 0
  }

  return limit - due
}

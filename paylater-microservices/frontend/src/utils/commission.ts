/**
 * Formats merchant commission DECIMAL strings for display (e.g. "5.00" → "5%").
 */
export function formatCommission(commission: string): string {
  const value = Number(commission)

  if (!Number.isFinite(value)) {
    return commission
  }

  return `${value}%`
}

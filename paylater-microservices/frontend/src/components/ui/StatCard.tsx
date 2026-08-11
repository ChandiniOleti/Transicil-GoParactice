export interface StatCardProps {
  label: string
  value: string
  className?: string
}

export default function StatCard({
  label,
  value,
  className = '',
}: StatCardProps) {
  return (
    <article className={`pl-stat-card ${className}`.trim()}>
      <p className="pl-stat-card__label">{label}</p>
      <p className="pl-stat-card__value">{value}</p>
    </article>
  )
}

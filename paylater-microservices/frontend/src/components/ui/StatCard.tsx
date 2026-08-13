export interface StatCardProps {
  label: string
  value: string
  className?: string
  clickable?: boolean
  onClick?: () => void
}

export default function StatCard({
  label,
  value,
  className = '',
  clickable = false,
  onClick,
}: StatCardProps) {
  const classes = [
    'pl-stat-card',
    clickable ? 'pl-stat-card--clickable' : '',
    className,
  ]
    .filter(Boolean)
    .join(' ')

  if (clickable && onClick) {
    return (
      <button
        type="button"
        className={classes}
        onClick={onClick}
        aria-label={`${label}: ${value}. View details.`}
      >
        <span className="pl-stat-card__label">{label}</span>
        <span className="pl-stat-card__value">{value}</span>
        <span className="pl-stat-card__hint">View details</span>
      </button>
    )
  }

  return (
    <article className={classes}>
      <p className="pl-stat-card__label">{label}</p>
      <p className="pl-stat-card__value">{value}</p>
    </article>
  )
}
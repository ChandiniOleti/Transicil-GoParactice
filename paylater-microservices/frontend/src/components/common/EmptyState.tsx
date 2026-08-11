import type { ReactNode } from 'react'

export interface EmptyStateProps {
  title: string
  description?: string
  action?: ReactNode
}

export default function EmptyState({
  title,
  description,
  action,
}: EmptyStateProps) {
  return (
    <div className="pl-empty">
      <h3 className="pl-empty__title">{title}</h3>
      {description ? <p className="pl-empty__description">{description}</p> : null}
      {action ? <div className="pl-empty__action">{action}</div> : null}
    </div>
  )
}

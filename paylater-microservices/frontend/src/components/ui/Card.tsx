import type { ReactNode } from 'react'

export interface CardProps {
  children: ReactNode
  title?: string
  className?: string
}

export default function Card({ children, title, className = '' }: CardProps) {
  return (
    <section className={`pl-card ${className}`.trim()}>
      {title ? <h2 className="pl-card__title">{title}</h2> : null}
      <div className="pl-card__body">{children}</div>
    </section>
  )
}

import type { ReactNode } from 'react'

export interface PageContainerProps {
  children: ReactNode
  title?: string
  description?: string
  className?: string
}

export default function PageContainer({
  children,
  title,
  description,
  className = '',
}: PageContainerProps) {
  return (
    <main className={`pl-page ${className}`.trim()}>
      {title || description ? (
        <header className="pl-page__header">
          {title ? <h1 className="pl-page__title">{title}</h1> : null}
          {description ? (
            <p className="pl-page__description">{description}</p>
          ) : null}
        </header>
      ) : null}
      <div className="pl-page__content">{children}</div>
    </main>
  )
}

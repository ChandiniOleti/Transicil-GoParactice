import type { ReactNode } from 'react'

export interface NavbarProps {
  brand?: string
  children?: ReactNode
  rightContent?: ReactNode
  className?: string
}

export default function Navbar({
  brand = 'PayLater',
  children,
  rightContent,
  className = '',
}: NavbarProps) {
  return (
    <header className={`pl-navbar ${className}`.trim()}>
      <div className="pl-navbar__brand">{brand}</div>
      {children ? <div className="pl-navbar__content">{children}</div> : null}
      {rightContent ? (
        <div className="pl-navbar__right">{rightContent}</div>
      ) : null}
    </header>
  )
}

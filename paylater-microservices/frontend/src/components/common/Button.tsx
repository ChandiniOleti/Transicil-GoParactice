import type { ButtonHTMLAttributes, ReactNode } from 'react'

type ButtonVariant = 'primary' | 'secondary' | 'danger' | 'outline'
type ButtonSize = 'small' | 'medium' | 'large'

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  children: ReactNode
  variant?: ButtonVariant
  size?: ButtonSize
  loading?: boolean
}

export default function Button({
  children,
  variant = 'primary',
  size = 'medium',
  type = 'button',
  disabled = false,
  loading = false,
  className = '',
  ...rest
}: ButtonProps) {
  const isDisabled = disabled || loading
  const classes = [
    'pl-button',
    `pl-button--${variant}`,
    `pl-button--${size}`,
    loading ? 'pl-button--loading' : '',
    className,
  ]
    .filter(Boolean)
    .join(' ')

  return (
    <button
      type={type}
      className={classes}
      disabled={isDisabled}
      aria-busy={loading || undefined}
      aria-disabled={isDisabled || undefined}
      {...rest}
    >
      {loading ? <span className="pl-button__spinner" aria-hidden="true" /> : null}
      <span className="pl-button__label">{loading ? 'Loading…' : children}</span>
    </button>
  )
}

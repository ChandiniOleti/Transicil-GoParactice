type LoaderSize = 'small' | 'medium' | 'large'

export interface LoaderProps {
  size?: LoaderSize
  label?: string
}

export default function Loader({
  size = 'medium',
  label = 'Loading',
}: LoaderProps) {
  return (
    <div
      className={`pl-loader pl-loader--${size}`}
      role="status"
      aria-live="polite"
      aria-label={label}
    >
      <span className="pl-loader__spinner" aria-hidden="true" />
      <span className="pl-loader__label">{label}</span>
    </div>
  )
}

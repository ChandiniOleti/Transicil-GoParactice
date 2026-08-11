import type { InputHTMLAttributes } from 'react'

export interface InputProps
  extends Omit<InputHTMLAttributes<HTMLInputElement>, 'size'> {
  label: string
  error?: string
  helperText?: string
}

export default function Input({
  label,
  id,
  name,
  type = 'text',
  error,
  helperText,
  className = '',
  disabled,
  required,
  ...rest
}: InputProps) {
  const inputId = id ?? name
  const errorId = error && inputId ? `${inputId}-error` : undefined
  const helperId =
    !error && helperText && inputId ? `${inputId}-helper` : undefined
  const describedBy = [errorId, helperId].filter(Boolean).join(' ') || undefined

  return (
    <div className={`pl-input ${className}`.trim()}>
      <label className="pl-input__label" htmlFor={inputId}>
        {label}
        {required ? <span aria-hidden="true"> *</span> : null}
      </label>
      <input
        id={inputId}
        name={name}
        type={type}
        className={`pl-input__field${error ? ' pl-input__field--error' : ''}`}
        disabled={disabled}
        required={required}
        aria-invalid={error ? true : undefined}
        aria-describedby={describedBy}
        {...rest}
      />
      {error ? (
        <p id={errorId} className="pl-input__error" role="alert">
          {error}
        </p>
      ) : null}
      {!error && helperText ? (
        <p id={helperId} className="pl-input__helper">
          {helperText}
        </p>
      ) : null}
    </div>
  )
}

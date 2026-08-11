import { useEffect, type ReactNode } from 'react'

export interface ModalProps {
  open: boolean
  title: string
  children: ReactNode
  onClose: () => void
  footer?: ReactNode
}

export default function Modal({
  open,
  title,
  children,
  onClose,
  footer,
}: ModalProps) {
  useEffect(() => {
    if (!open) {
      return
    }

    const onKeyDown = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        onClose()
      }
    }

    document.addEventListener('keydown', onKeyDown)
    return () => {
      document.removeEventListener('keydown', onKeyDown)
    }
  }, [open, onClose])

  if (!open) {
    return null
  }

  return (
    <div className="pl-modal" role="presentation">
      <button
        type="button"
        className="pl-modal__backdrop"
        aria-label="Close dialog"
        onClick={onClose}
      />
      <div
        className="pl-modal__dialog"
        role="dialog"
        aria-modal="true"
        aria-labelledby="pl-modal-title"
      >
        <header className="pl-modal__header">
          <h2 id="pl-modal-title" className="pl-modal__title">
            {title}
          </h2>
          <button
            type="button"
            className="pl-modal__close"
            aria-label="Close"
            onClick={onClose}
          >
            ×
          </button>
        </header>
        <div className="pl-modal__body">{children}</div>
        {footer ? <footer className="pl-modal__footer">{footer}</footer> : null}
      </div>
    </div>
  )
}

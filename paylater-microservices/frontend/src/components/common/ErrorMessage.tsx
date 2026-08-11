export interface ErrorMessageProps {
  message: string
  title?: string
}

export default function ErrorMessage({
  message,
  title = 'Error',
}: ErrorMessageProps) {
  return (
    <div className="pl-error" role="alert">
      <strong className="pl-error__title">{title}</strong>
      <p className="pl-error__message">{message}</p>
    </div>
  )
}

import { Link } from 'react-router-dom'

export default function NotFoundPage() {
  return (
    <div className="pl-not-found">
      <h1 className="pl-not-found__title">Page Not Found</h1>
      <p className="pl-not-found__message">
        The page you requested does not exist.
      </p>
      <Link className="pl-not-found__link" to="/login">
        Back to Login
      </Link>
    </div>
  )
}

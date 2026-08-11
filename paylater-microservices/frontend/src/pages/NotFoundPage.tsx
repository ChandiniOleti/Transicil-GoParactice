import { Link } from 'react-router-dom'

export default function NotFoundPage() {
  return (
    <div>
      <h1>Page Not Found</h1>
      <p>The page you requested does not exist.</p>
      <Link to="/login">Back to Login</Link>
    </div>
  )
}

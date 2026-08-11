import Card from '../../components/ui/Card'
import type { User } from '../../types/user'

export interface UserSummaryProps {
  user: User
}

export default function UserSummary({ user }: UserSummaryProps) {
  return (
    <Card title="Account" className="pl-user-summary">
      <dl className="pl-user-summary__list">
        <div>
          <dt>Name</dt>
          <dd>{user.name}</dd>
        </div>
        <div>
          <dt>Email</dt>
          <dd>{user.email}</dd>
        </div>
      </dl>
    </Card>
  )
}

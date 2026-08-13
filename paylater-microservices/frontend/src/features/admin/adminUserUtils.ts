import type { User } from '../../types/user'

function parseAmount(value: string): number {
  const parsed = Number(value)
  return Number.isFinite(parsed) ? parsed : 0
}

export function sumUserTotalDue(users: User[]): number {
  return users.reduce((sum, user) => sum + parseAmount(user.current_due), 0)
}

export function filterUsersByName(users: User[], query: string): User[] {
  const normalizedQuery = query.trim().toLowerCase()

  if (!normalizedQuery) {
    return []
  }

  return users.filter((user) =>
    user.name.toLowerCase().includes(normalizedQuery),
  )
}

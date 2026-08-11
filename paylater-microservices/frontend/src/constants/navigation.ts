import type { NavigationItem } from '../types/navigation'

export const userNavigation: NavigationItem[] = [
  { label: 'Dashboard', path: '/user/dashboard' },
  { label: 'Profile', path: '/user/profile' },
  { label: 'Transactions', path: '/user/transactions' },
  { label: 'Payback', path: '/user/payback' },
]

export const merchantNavigation: NavigationItem[] = [
  { label: 'Dashboard', path: '/merchant/dashboard' },
  { label: 'Profile', path: '/merchant/profile' },
  { label: 'Transactions', path: '/merchant/transactions' },
]

export const adminNavigation: NavigationItem[] = [
  { label: 'Dashboard', path: '/admin/dashboard' },
  { label: 'Users', path: '/admin/users' },
  { label: 'Merchants', path: '/admin/merchants' },
  { label: 'Transactions', path: '/admin/transactions' },
  { label: 'Reports', path: '/admin/reports' },
  { label: 'Create Admin', path: '/admin/create-admin' },
]

import type { NavigationItem } from '../types/navigation'

export const userNavigation: NavigationItem[] = [
  { label: 'Dashboard', path: '/user/dashboard' },
  { label: 'Profile', path: '/user/profile' },
  { label: 'Purchase', path: '/user/purchase' },
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
  { label: 'Create User', path: '/admin/create-user' },
  { label: 'Merchants', path: '/admin/merchants' },
  { label: 'Create Merchant', path: '/admin/create-merchant' },
  { label: 'Transactions', path: '/admin/transactions' },
  { label: 'Reports', path: '/admin/reports' },
  { label: 'All Admins', path: '/admin/admins' },
  { label: 'Create Admin', path: '/admin/create-admin' },
]

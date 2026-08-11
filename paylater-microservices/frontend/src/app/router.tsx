import {
  createBrowserRouter,
  Navigate,
  RouterProvider,
} from 'react-router-dom'

import AppLayout from '../components/layout/AppLayout'
import PublicLayout from '../components/layout/PublicLayout'
import {
  adminNavigation,
  merchantNavigation,
  userNavigation,
} from '../constants/navigation'
import LoginPage from '../pages/auth/LoginPage'
import RegisterPage from '../pages/auth/RegisterPage'
import UserDashboardPage from '../pages/user/UserDashboardPage'
import UserProfilePage from '../pages/user/UserProfilePage'
import UserTransactionsPage from '../pages/user/UserTransactionsPage'
import UserPaybackPage from '../pages/user/UserPaybackPage'
import MerchantDashboardPage from '../pages/merchant/MerchantDashboardPage'
import MerchantProfilePage from '../pages/merchant/MerchantProfilePage'
import MerchantTransactionsPage from '../pages/merchant/MerchantTransactionsPage'
import AdminDashboardPage from '../pages/admin/AdminDashboardPage'
import AdminUsersPage from '../pages/admin/AdminUsersPage'
import AdminMerchantsPage from '../pages/admin/AdminMerchantsPage'
import AdminTransactionsPage from '../pages/admin/AdminTransactionsPage'
import AdminReportsPage from '../pages/admin/AdminReportsPage'
import AdminCreateAdminPage from '../pages/admin/AdminCreateAdminPage'
import NotFoundPage from '../pages/NotFoundPage'
import GuestRoute from '../routes/GuestRoute'
import ProtectedRoute from '../routes/ProtectedRoute'
import RoleRoute from '../routes/RoleRoute'

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Navigate to="/login" replace />,
  },
  {
    element: <PublicLayout />,
    children: [
      {
        element: <GuestRoute />,
        children: [
          { path: '/login', element: <LoginPage /> },
          { path: '/register', element: <RegisterPage /> },
        ],
      },
    ],
  },
  {
    element: <ProtectedRoute />,
    children: [
      {
        element: <RoleRoute allowedRoles={['USER']} />,
        children: [
          {
            path: '/user',
            element: <AppLayout navigationItems={userNavigation} />,
            children: [
              { index: true, element: <Navigate to="dashboard" replace /> },
              { path: 'dashboard', element: <UserDashboardPage /> },
              { path: 'profile', element: <UserProfilePage /> },
              { path: 'transactions', element: <UserTransactionsPage /> },
              { path: 'payback', element: <UserPaybackPage /> },
            ],
          },
        ],
      },
      {
        element: <RoleRoute allowedRoles={['MERCHANT']} />,
        children: [
          {
            path: '/merchant',
            element: <AppLayout navigationItems={merchantNavigation} />,
            children: [
              { index: true, element: <Navigate to="dashboard" replace /> },
              { path: 'dashboard', element: <MerchantDashboardPage /> },
              { path: 'profile', element: <MerchantProfilePage /> },
              { path: 'transactions', element: <MerchantTransactionsPage /> },
            ],
          },
        ],
      },
      {
        element: <RoleRoute allowedRoles={['ADMIN']} />,
        children: [
          {
            path: '/admin',
            element: <AppLayout navigationItems={adminNavigation} />,
            children: [
              { index: true, element: <Navigate to="dashboard" replace /> },
              { path: 'dashboard', element: <AdminDashboardPage /> },
              { path: 'users', element: <AdminUsersPage /> },
              { path: 'merchants', element: <AdminMerchantsPage /> },
              { path: 'transactions', element: <AdminTransactionsPage /> },
              { path: 'reports', element: <AdminReportsPage /> },
              { path: 'create-admin', element: <AdminCreateAdminPage /> },
            ],
          },
        ],
      },
    ],
  },
  {
    path: '*',
    element: <NotFoundPage />,
  },
])

export default function AppRouter() {
  return <RouterProvider router={router} />
}

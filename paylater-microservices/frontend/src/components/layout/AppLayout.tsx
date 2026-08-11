import { Outlet } from 'react-router-dom'

import type { NavigationItem } from '../../types/navigation'
import AccountMenu from './AccountMenu'
import Navbar from './Navbar'
import Sidebar from './Sidebar'

export interface AppLayoutProps {
  navigationItems: NavigationItem[]
}

/**
 * Authenticated application shell.
 * Navigation is injected via Props so the layout stays role-agnostic.
 */
export default function AppLayout({ navigationItems }: AppLayoutProps) {
  return (
    <div className="pl-app-layout">
      <Navbar brand="PayLater" rightContent={<AccountMenu />} />
      <div className="pl-app-layout__body">
        <Sidebar items={navigationItems} />
        <main className="pl-app-layout__main" id="main-content">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

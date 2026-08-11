import { Outlet } from 'react-router-dom'

/**
 * Public application shell for unauthenticated pages (login/register).
 * No sidebar or authenticated navigation.
 */
export default function PublicLayout() {
  return (
    <div className="pl-public-layout">
      <header className="pl-public-layout__brand">PayLater</header>
      <main className="pl-public-layout__main" id="main-content">
        <Outlet />
      </main>
    </div>
  )
}

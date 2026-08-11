import { NavLink } from 'react-router-dom'

import type { NavigationItem } from '../../types/navigation'

export type { NavigationItem }

export interface SidebarProps {
  items: NavigationItem[]
  className?: string
}

export default function Sidebar({ items, className = '' }: SidebarProps) {
  return (
    <aside className={`pl-sidebar ${className}`.trim()} aria-label="Sidebar">
      <nav aria-label="Primary">
        <ul className="pl-sidebar__list">
          {items.map((item) => (
            <li key={item.path} className="pl-sidebar__item">
              <NavLink
                to={item.path}
                className={({ isActive }) =>
                  [
                    'pl-sidebar__link',
                    isActive ? 'pl-sidebar__link--active' : '',
                  ]
                    .filter(Boolean)
                    .join(' ')
                }
              >
                {item.icon ? (
                  <span className="pl-sidebar__icon" aria-hidden="true">
                    {item.icon}
                  </span>
                ) : null}
                <span>{item.label}</span>
              </NavLink>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  )
}

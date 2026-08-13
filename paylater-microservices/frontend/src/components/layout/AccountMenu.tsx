import { useEffect, useRef, useState } from 'react'
import { useNavigate } from 'react-router-dom'

import { useAppDispatch, useAppSelector } from '../../app/hooks'
import { logout } from '../../features/auth/authSlice'
import {
  selectAuthRole,
  selectCurrentUser,
} from '../../features/auth/authSelectors'

export default function AccountMenu() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()
  const user = useAppSelector(selectCurrentUser)
  const role = useAppSelector(selectAuthRole)
  const [isOpen, setIsOpen] = useState(false)
  const menuRef = useRef<HTMLDivElement>(null)

  const displayLabel = user?.email ?? 'Account'

  useEffect(() => {
    if (!isOpen) {
      return
    }

    function handleClickOutside(event: MouseEvent) {
      if (
        menuRef.current &&
        !menuRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false)
      }
    }

    function handleEscape(event: KeyboardEvent) {
      if (event.key === 'Escape') {
        setIsOpen(false)
      }
    }

    document.addEventListener('mousedown', handleClickOutside)
    document.addEventListener('keydown', handleEscape)

    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
      document.removeEventListener('keydown', handleEscape)
    }
  }, [isOpen])

  function handleLogout() {
    dispatch(logout())
    setIsOpen(false)
    navigate('/login', { replace: true })
  }

  return (
    <div className="pl-account-menu" ref={menuRef}>
      <button
        type="button"
        className="pl-account-menu__trigger"
        aria-haspopup="menu"
        aria-expanded={isOpen}
        onClick={() => setIsOpen((open) => !open)}
      >
        <span className="pl-account-menu__label">{displayLabel}</span>
        <span className="pl-account-menu__chevron" aria-hidden="true">
          ▾
        </span>
      </button>

      {isOpen ? (
        <div className="pl-account-menu__dropdown" role="menu">
          <div className="pl-account-menu__info">
            {user?.email ? (
              <p className="pl-account-menu__email">{user.email}</p>
            ) : null}
            {role ? (
              <p className="pl-account-menu__role">{role}</p>
            ) : null}
          </div>
          <button
            type="button"
            className="pl-account-menu__logout"
            role="menuitem"
            onClick={handleLogout}
          >
            Logout
          </button>
        </div>
      ) : null}
    </div>
  )
}

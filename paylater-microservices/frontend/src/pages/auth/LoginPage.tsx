import { useState, type FormEvent } from 'react'
import { Link, useNavigate } from 'react-router-dom'

import { useAppDispatch } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import Card from '../../components/ui/Card'
import { setCredentials, type AuthUser } from '../../features/auth/authSlice'
import { adminLogin, login, merchantLogin } from '../../services/authApi'
import type { AuthRole, LoginResponse } from '../../types/auth'
import { getErrorMessage } from '../../utils/error'
import { readJwtClaims } from '../../utils/jwt'
import { getDashboardPath } from '../../utils/routes'

interface FormErrors {
  email?: string
  password?: string
}

const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

const LOGIN_OPTIONS: Array<{ value: AuthRole; label: string }> = [
  { value: 'USER', label: 'User' },
  { value: 'ADMIN', label: 'Admin' },
  { value: 'MERCHANT', label: 'Merchant' },
]

async function authenticate(
  role: AuthRole,
  email: string,
  password: string,
): Promise<LoginResponse> {
  switch (role) {
    case 'ADMIN':
      return adminLogin(email, password)
    case 'MERCHANT':
      return merchantLogin(email, password)
    case 'USER':
      return login(email, password)
  }
}

function buildAuthUser(
  token: string,
  fallbackEmail: string,
  fallbackRole: AuthRole,
): AuthUser {
  const claims = readJwtClaims(token)
  if (!claims) {
    throw new Error('Login succeeded but authentication details were invalid.')
  }

  return {
    user_id: claims.user_id,
    email: claims.email || fallbackEmail,
    role: claims.role || fallbackRole,
  }
}

export default function LoginPage() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const [loginRole, setLoginRole] = useState<AuthRole>('USER')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isPasswordVisible, setIsPasswordVisible] = useState(false)
  const [fieldErrors, setFieldErrors] = useState<FormErrors>({})
  const [formError, setFormError] = useState<string | null>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)

  function validate(): FormErrors {
    const errors: FormErrors = {}
    const trimmedEmail = email.trim()

    if (!trimmedEmail) {
      errors.email = 'Email is required.'
    } else if (!EMAIL_PATTERN.test(trimmedEmail)) {
      errors.email = 'Enter a valid email address.'
    }

    if (!password) {
      errors.password = 'Password is required.'
    }

    return errors
  }

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    if (isSubmitting) {
      return
    }

    setFormError(null)
    const errors = validate()
    setFieldErrors(errors)
    if (Object.keys(errors).length > 0) {
      return
    }

    const trimmedEmail = email.trim()

    setIsSubmitting(true)
    try {
      const response = await authenticate(loginRole, trimmedEmail, password)
      const user = buildAuthUser(response.token, trimmedEmail, loginRole)

      dispatch(
        setCredentials({
          token: response.token,
          user,
        }),
      )

      navigate(getDashboardPath(user.role), { replace: true })
    } catch (error) {
      setFormError(getErrorMessage(error))
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <Card title="Sign In" className="pl-login-card">
      <form className="pl-login-form" onSubmit={handleSubmit} noValidate>
        <fieldset className="pl-login-role" disabled={isSubmitting}>
          <legend className="pl-login-role__legend">Sign in as</legend>
          <div
            className="pl-login-role__options"
            role="radiogroup"
            aria-label="Login type"
          >
            {LOGIN_OPTIONS.map((option) => (
              <label key={option.value} className="pl-login-role__option">
                <input
                  type="radio"
                  name="loginRole"
                  value={option.value}
                  checked={loginRole === option.value}
                  onChange={() => setLoginRole(option.value)}
                />
                <span>{option.label}</span>
              </label>
            ))}
          </div>
        </fieldset>

        <Input
          label="Email"
          id="login-email"
          name="email"
          type="email"
          autoComplete="email"
          value={email}
          onChange={(event) => setEmail(event.target.value)}
          error={fieldErrors.email}
          disabled={isSubmitting}
          required
        />

        <div className="pl-input pl-login-password">
          <label className="pl-input__label" htmlFor="login-password">
            Password<span aria-hidden="true"> *</span>
          </label>
          <div className="pl-login-password__control">
            <input
              id="login-password"
              name="password"
              type={isPasswordVisible ? 'text' : 'password'}
              autoComplete="current-password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
              className={`pl-input__field pl-login-password__field${
                fieldErrors.password ? ' pl-input__field--error' : ''
              }`}
              disabled={isSubmitting}
              required
              aria-invalid={fieldErrors.password ? true : undefined}
              aria-describedby={
                fieldErrors.password ? 'login-password-error' : undefined
              }
            />
            <button
              type="button"
              className="pl-login-password__toggle"
              onClick={() => setIsPasswordVisible((visible) => !visible)}
              disabled={isSubmitting}
              aria-label={isPasswordVisible ? 'Hide password' : 'Show password'}
              aria-pressed={isPasswordVisible}
            >
              <span aria-hidden="true">{isPasswordVisible ? '🙈' : '👁'}</span>
            </button>
          </div>
          {fieldErrors.password ? (
            <p id="login-password-error" className="pl-input__error" role="alert">
              {fieldErrors.password}
            </p>
          ) : null}
        </div>

        {formError ? <ErrorMessage message={formError} title="Login failed" /> : null}

        <Button type="submit" variant="primary" loading={isSubmitting} className="pl-login-form__submit">
          Login
        </Button>

        <p className="pl-login-form__footer">
          <Link to="/register">Create an account</Link>
        </p>
      </form>
    </Card>
  )
}

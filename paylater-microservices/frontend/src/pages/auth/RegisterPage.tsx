import { useState, type FormEvent } from 'react'
import { Link, useNavigate } from 'react-router-dom'

import { useAppDispatch } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import Card from '../../components/ui/Card'
import { setCredentials, type AuthUser } from '../../features/auth/authSlice'
import { login, merchantLogin } from '../../services/authApi'
import { createMerchant } from '../../services/merchantApi'
import { createUser } from '../../services/userApi'
import type { AuthRole } from '../../types/auth'
import { getErrorMessage } from '../../utils/error'
import { readJwtClaims } from '../../utils/jwt'

type RegisterRole = Extract<AuthRole, 'USER' | 'MERCHANT'>

interface FormErrors {
  name?: string
  merchantName?: string
  email?: string
  phone?: string
  commission?: string
  password?: string
  confirmPassword?: string
}

const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

const REGISTER_OPTIONS: Array<{ value: RegisterRole; label: string }> = [
  { value: 'USER', label: 'User' },
  { value: 'MERCHANT', label: 'Merchant' },
]

function buildAuthUser(
  token: string,
  fallbackEmail: string,
  fallbackRole: RegisterRole,
): AuthUser {
  const claims = readJwtClaims(token)
  if (!claims) {
    throw new Error('Registration succeeded but authentication details were invalid.')
  }

  return {
    user_id: claims.user_id,
    email: claims.email || fallbackEmail,
    role: claims.role || fallbackRole,
  }
}

function validateEmail(email: string): string | undefined {
  const trimmedEmail = email.trim()

  if (!trimmedEmail) {
    return 'Email is required.'
  }

  if (!EMAIL_PATTERN.test(trimmedEmail)) {
    return 'Enter a valid email address.'
  }

  return undefined
}

function validatePasswordFields(
  password: string,
  confirmPassword: string,
): Pick<FormErrors, 'password' | 'confirmPassword'> {
  const errors: Pick<FormErrors, 'password' | 'confirmPassword'> = {}

  if (!password) {
    errors.password = 'Password is required.'
  }

  if (!confirmPassword) {
    errors.confirmPassword = 'Please confirm your password.'
  } else if (password !== confirmPassword) {
    errors.confirmPassword = 'Passwords do not match.'
  }

  return errors
}

function validateCommission(commission: string): string | undefined {
  const trimmed = commission.trim()

  if (!trimmed) {
    return 'Commission is required.'
  }

  const value = Number(trimmed)
  if (!Number.isFinite(value)) {
    return 'Commission must be a number.'
  }

  if (value < 3) {
    return 'Commission must be at least 3%.'
  }

  if (value > 10) {
    return 'Commission must not exceed 10%.'
  }

  return undefined
}

export default function RegisterPage() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const [registerRole, setRegisterRole] = useState<RegisterRole>('USER')
  const [name, setName] = useState('')
  const [merchantName, setMerchantName] = useState('')
  const [email, setEmail] = useState('')
  const [phone, setPhone] = useState('')
  const [commission, setCommission] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [fieldErrors, setFieldErrors] = useState<FormErrors>({})
  const [formError, setFormError] = useState<string | null>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)

  function validateUserForm(): FormErrors {
    const errors: FormErrors = {}
    const trimmedName = name.trim()

    if (!trimmedName) {
      errors.name = 'Name is required.'
    }

    const emailError = validateEmail(email)
    if (emailError) {
      errors.email = emailError
    }

    Object.assign(errors, validatePasswordFields(password, confirmPassword))

    return errors
  }

  function validateMerchantForm(): FormErrors {
    const errors: FormErrors = {}
    const trimmedMerchantName = merchantName.trim()
    const trimmedPhone = phone.trim()

    if (!trimmedMerchantName) {
      errors.merchantName = 'Merchant name is required.'
    }

    const emailError = validateEmail(email)
    if (emailError) {
      errors.email = emailError
    }

    if (!trimmedPhone) {
      errors.phone = 'Phone is required.'
    }

    const commissionError = validateCommission(commission)
    if (commissionError) {
      errors.commission = commissionError
    }

    Object.assign(errors, validatePasswordFields(password, confirmPassword))

    return errors
  }

  function handleRoleChange(role: RegisterRole) {
    setRegisterRole(role)
    setFieldErrors({})
    setFormError(null)
  }

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    if (isSubmitting) {
      return
    }

    setFormError(null)
    const errors =
      registerRole === 'USER' ? validateUserForm() : validateMerchantForm()
    setFieldErrors(errors)
    if (Object.keys(errors).length > 0) {
      return
    }

    const trimmedEmail = email.trim()

    setIsSubmitting(true)
    try {
      if (registerRole === 'USER') {
        await createUser({
          name: name.trim(),
          email: trimmedEmail,
          password,
        })

        const loginResponse = await login(trimmedEmail, password)
        const user = buildAuthUser(loginResponse.token, trimmedEmail, 'USER')

        dispatch(
          setCredentials({
            token: loginResponse.token,
            user,
          }),
        )

        navigate('/user/dashboard', { replace: true })
        return
      }

      await createMerchant({
        merchant_name: merchantName.trim(),
        phone: phone.trim(),
        commission: commission.trim(),
        email: trimmedEmail,
        password,
      })

      const loginResponse = await merchantLogin(trimmedEmail, password)
      const user = buildAuthUser(loginResponse.token, trimmedEmail, 'MERCHANT')

      dispatch(
        setCredentials({
          token: loginResponse.token,
          user,
        }),
      )

      navigate('/merchant/dashboard', { replace: true })
    } catch (error) {
      setFormError(getErrorMessage(error))
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <Card title="Create Account" className="pl-login-card">
      <form className="pl-login-form" onSubmit={handleSubmit} noValidate>
        <fieldset className="pl-login-role" disabled={isSubmitting}>
          <legend className="pl-login-role__legend">Register as</legend>
          <div
            className="pl-login-role__options"
            role="radiogroup"
            aria-label="Registration type"
          >
            {REGISTER_OPTIONS.map((option) => (
              <label key={option.value} className="pl-login-role__option">
                <input
                  type="radio"
                  name="registerRole"
                  value={option.value}
                  checked={registerRole === option.value}
                  onChange={() => handleRoleChange(option.value)}
                />
                <span>{option.label}</span>
              </label>
            ))}
          </div>
        </fieldset>

        {registerRole === 'USER' ? (
          <Input
            label="Full Name"
            id="register-name"
            name="name"
            type="text"
            autoComplete="name"
            value={name}
            onChange={(event) => setName(event.target.value)}
            error={fieldErrors.name}
            disabled={isSubmitting}
            required
          />
        ) : (
          <>
            <Input
              label="Merchant Name"
              id="register-merchant-name"
              name="merchantName"
              type="text"
              autoComplete="organization"
              value={merchantName}
              onChange={(event) => setMerchantName(event.target.value)}
              error={fieldErrors.merchantName}
              disabled={isSubmitting}
              required
            />

            <Input
              label="Phone"
              id="register-phone"
              name="phone"
              type="tel"
              autoComplete="tel"
              value={phone}
              onChange={(event) => setPhone(event.target.value)}
              error={fieldErrors.phone}
              disabled={isSubmitting}
              required
            />

            <Input
              label="Commission (%)"
              id="register-commission"
              name="commission"
              type="number"
              inputMode="decimal"
              step="0.01"
              min="3"
              max="10"
              value={commission}
              onChange={(event) => setCommission(event.target.value)}
              error={fieldErrors.commission}
              helperText="Must be between 3% and 10%."
              disabled={isSubmitting}
              required
            />
          </>
        )}

        <Input
          label="Email"
          id="register-email"
          name="email"
          type="email"
          autoComplete="email"
          value={email}
          onChange={(event) => setEmail(event.target.value)}
          error={fieldErrors.email}
          disabled={isSubmitting}
          required
        />

        <Input
          label="Password"
          id="register-password"
          name="password"
          type="password"
          autoComplete="new-password"
          value={password}
          onChange={(event) => setPassword(event.target.value)}
          error={fieldErrors.password}
          disabled={isSubmitting}
          required
        />

        <Input
          label="Confirm Password"
          id="register-confirm-password"
          name="confirmPassword"
          type="password"
          autoComplete="new-password"
          value={confirmPassword}
          onChange={(event) => setConfirmPassword(event.target.value)}
          error={fieldErrors.confirmPassword}
          disabled={isSubmitting}
          required
        />

        {formError ? (
          <ErrorMessage message={formError} title="Registration failed" />
        ) : null}

        <Button
          type="submit"
          variant="primary"
          loading={isSubmitting}
          className="pl-login-form__submit"
        >
          Register
        </Button>

        <p className="pl-login-form__footer">
          <Link to="/login">Already have an account? Sign in</Link>
        </p>
      </form>
    </Card>
  )
}

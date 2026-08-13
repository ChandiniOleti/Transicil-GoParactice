import { useState, type FormEvent } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import { createMerchant } from '../../services/merchantApi'
import { getErrorMessage } from '../../utils/error'

interface FormErrors {
  merchantName?: string
  email?: string
  phone?: string
  commission?: string
  password?: string
  confirmPassword?: string
}

const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

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

export default function AdminCreateMerchantPage() {
  const [merchantName, setMerchantName] = useState('')
  const [email, setEmail] = useState('')
  const [phone, setPhone] = useState('')
  const [commission, setCommission] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [fieldErrors, setFieldErrors] = useState<FormErrors>({})
  const [submitError, setSubmitError] = useState<string | null>(null)
  const [successMessage, setSuccessMessage] = useState<string | null>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)

  function validate(): FormErrors {
    const errors: FormErrors = {}
    const trimmedMerchantName = merchantName.trim()
    const trimmedEmail = email.trim()
    const trimmedPhone = phone.trim()

    if (!trimmedMerchantName) {
      errors.merchantName = 'Merchant name is required.'
    }

    if (!trimmedEmail) {
      errors.email = 'Email is required.'
    } else if (!EMAIL_PATTERN.test(trimmedEmail)) {
      errors.email = 'Enter a valid email address.'
    }

    if (!trimmedPhone) {
      errors.phone = 'Phone is required.'
    }

    const commissionError = validateCommission(commission)
    if (commissionError) {
      errors.commission = commissionError
    }

    if (!password) {
      errors.password = 'Password is required.'
    }

    if (!confirmPassword) {
      errors.confirmPassword = 'Please confirm the password.'
    } else if (password !== confirmPassword) {
      errors.confirmPassword = 'Passwords do not match.'
    }

    return errors
  }

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    if (isSubmitting) {
      return
    }

    setSubmitError(null)
    setSuccessMessage(null)

    const errors = validate()
    setFieldErrors(errors)
    if (Object.keys(errors).length > 0) {
      return
    }

    setIsSubmitting(true)
    try {
      const response = await createMerchant({
        merchant_name: merchantName.trim(),
        phone: phone.trim(),
        commission: commission.trim(),
        email: email.trim(),
        password,
      })

      setSuccessMessage(response.message)
      setMerchantName('')
      setEmail('')
      setPhone('')
      setCommission('')
      setPassword('')
      setConfirmPassword('')
      setFieldErrors({})
    } catch (error) {
      setSubmitError(getErrorMessage(error))
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <PageContainer
      title="Create Merchant"
      description="Add a new PayLater merchant account."
    >
      <div className="pl-admin-create-form">
        <Card title="New Merchant Account">
          <form
            className="pl-admin-create-form__form"
            onSubmit={(event) => void handleSubmit(event)}
            noValidate
          >
            {successMessage ? (
              <div className="pl-user-payback__success" role="status">
                <strong className="pl-user-payback__success-title">
                  Success
                </strong>
                <p className="pl-user-payback__success-message">
                  {successMessage}
                </p>
              </div>
            ) : null}

            {submitError ? (
              <ErrorMessage
                title="Unable to create merchant"
                message={submitError}
              />
            ) : null}

            <Input
              label="Merchant Name"
              id="create-merchant-name"
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
              label="Email"
              id="create-merchant-email"
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
              label="Phone"
              id="create-merchant-phone"
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
              id="create-merchant-commission"
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

            <Input
              label="Password"
              id="create-merchant-password"
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
              id="create-merchant-confirm-password"
              name="confirmPassword"
              type="password"
              autoComplete="new-password"
              value={confirmPassword}
              onChange={(event) => setConfirmPassword(event.target.value)}
              error={fieldErrors.confirmPassword}
              disabled={isSubmitting}
              required
            />

            <Button type="submit" loading={isSubmitting} disabled={isSubmitting}>
              Create Merchant
            </Button>
          </form>
        </Card>
      </div>
    </PageContainer>
  )
}

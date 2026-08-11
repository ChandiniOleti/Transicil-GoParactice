import { useCallback, useEffect, useState, type FormEvent } from 'react'

import { useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import StatCard from '../../components/ui/StatCard'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import { createPayback } from '../../services/paybackApi'
import { getUserById } from '../../services/userApi'
import type { PaybackResponse } from '../../types/payback'
import type { User } from '../../types/user'
import { formatCurrency } from '../../utils/currency'
import { getErrorMessage } from '../../utils/error'

interface FormErrors {
  amount?: string
}

function validateAmount(amount: string, currentDue: string): FormErrors {
  const trimmed = amount.trim()

  if (!trimmed) {
    return { amount: 'Amount is required.' }
  }

  const parsedAmount = Number(trimmed)
  if (!Number.isFinite(parsedAmount) || parsedAmount <= 0) {
    return { amount: 'Amount must be greater than 0.' }
  }

  const due = Number(currentDue)
  if (Number.isFinite(due) && parsedAmount > due) {
    return { amount: 'Amount must not exceed current due.' }
  }

  return {}
}

export default function UserPaybackPage() {
  const authUser = useAppSelector(selectCurrentUser)
  const [user, setUser] = useState<User | null>(null)
  const [isLoadingUser, setIsLoadingUser] = useState(false)
  const [loadError, setLoadError] = useState<string | null>(null)

  const [amount, setAmount] = useState('')
  const [fieldErrors, setFieldErrors] = useState<FormErrors>({})
  const [submitError, setSubmitError] = useState<string | null>(null)
  const [successMessage, setSuccessMessage] = useState<string | null>(null)
  const [lastPayback, setLastPayback] = useState<PaybackResponse | null>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)

  const loadUser = useCallback(async () => {
    if (!authUser) {
      setUser(null)
      setLoadError('Authenticated user information is unavailable.')
      return
    }

    setIsLoadingUser(true)
    setLoadError(null)

    try {
      const response = await getUserById(authUser.user_id)
      setUser(response)
    } catch (err) {
      setUser(null)
      setLoadError(getErrorMessage(err))
    } finally {
      setIsLoadingUser(false)
    }
  }, [authUser])

  useEffect(() => {
    void loadUser()
  }, [loadUser])

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    if (!authUser || !user || isSubmitting) {
      return
    }

    setSubmitError(null)
    setSuccessMessage(null)
    setLastPayback(null)

    const errors = validateAmount(amount, user.current_due)
    setFieldErrors(errors)

    if (errors.amount) {
      return
    }

    setIsSubmitting(true)

    try {
      const response = await createPayback({
        user_id: authUser.user_id,
        amount: amount.trim(),
      })

      setSuccessMessage(response.message)
      setLastPayback(response)
      setAmount('')
      setFieldErrors({})
      await loadUser()
    } catch (err) {
      setSubmitError(getErrorMessage(err))
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <PageContainer
      title="Payback"
      description="Pay down your current balance."
    >
      {isLoadingUser ? <Loader label="Loading account details" /> : null}

      {!isLoadingUser && loadError ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load account" message={loadError} />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadUser()}
          >
            Retry
          </Button>
        </div>
      ) : null}

      {!isLoadingUser && !loadError && user ? (
        <div className="pl-user-payback">
          <StatCard
            label="Current Due"
            value={formatCurrency(user.current_due)}
          />

          <Card title="Make a Payback">
            <form
              className="pl-user-payback__form"
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
                  {lastPayback ? (
                    <p className="pl-user-payback__success-message">
                      Amount paid: {formatCurrency(lastPayback.amount_paid)} ·
                      Updated current due:{' '}
                      {formatCurrency(lastPayback.updated_current_due)}
                    </p>
                  ) : null}
                </div>
              ) : null}

              {submitError ? (
                <ErrorMessage title="Payback failed" message={submitError} />
              ) : null}

              <Input
                label="Amount"
                name="amount"
                type="number"
                inputMode="decimal"
                step="0.01"
                min="0"
                value={amount}
                onChange={(event) => setAmount(event.target.value)}
                error={fieldErrors.amount}
                helperText={`Maximum payback: ${formatCurrency(user.current_due)}`}
                disabled={isSubmitting}
                required
              />

              <Button type="submit" loading={isSubmitting} disabled={isSubmitting}>
                Payback
              </Button>
            </form>
          </Card>
        </div>
      ) : null}
    </PageContainer>
  )
}

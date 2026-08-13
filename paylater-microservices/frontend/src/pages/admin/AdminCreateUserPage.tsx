import { useState, type FormEvent } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import { createUser } from '../../services/userApi'
import { getErrorMessage } from '../../utils/error'

interface FormErrors {
  name?: string
  email?: string
  password?: string
  confirmPassword?: string
}

const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

export default function AdminCreateUserPage() {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [fieldErrors, setFieldErrors] = useState<FormErrors>({})
  const [submitError, setSubmitError] = useState<string | null>(null)
  const [successMessage, setSuccessMessage] = useState<string | null>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)

  function validate(): FormErrors {
    const errors: FormErrors = {}
    const trimmedName = name.trim()
    const trimmedEmail = email.trim()

    if (!trimmedName) {
      errors.name = 'Name is required.'
    }

    if (!trimmedEmail) {
      errors.email = 'Email is required.'
    } else if (!EMAIL_PATTERN.test(trimmedEmail)) {
      errors.email = 'Enter a valid email address.'
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
      const response = await createUser({
        name: name.trim(),
        email: email.trim(),
        password,
      })

      setSuccessMessage(response.message)
      setName('')
      setEmail('')
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
      title="Create User"
      description="Add a new PayLater customer account."
    >
      <div className="pl-admin-create-form">
        <Card title="New User Account">
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
              <ErrorMessage title="Unable to create user" message={submitError} />
            ) : null}

            <Input
              label="Name"
              id="create-user-name"
              name="name"
              type="text"
              autoComplete="name"
              value={name}
              onChange={(event) => setName(event.target.value)}
              error={fieldErrors.name}
              disabled={isSubmitting}
              required
            />

            <Input
              label="Email"
              id="create-user-email"
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
              id="create-user-password"
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
              id="create-user-confirm-password"
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
              Create User
            </Button>
          </form>
        </Card>
      </div>
    </PageContainer>
  )
}

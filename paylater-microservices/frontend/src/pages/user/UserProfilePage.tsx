import { useCallback, useEffect, useState, type FormEvent } from 'react'
import { useNavigate } from 'react-router-dom'

import { useAppDispatch, useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import Modal from '../../components/ui/Modal'
import ProfileField from '../../components/ui/ProfileField'
import { logout } from '../../features/auth/authSlice'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import { deleteUser, getUserById, updateUser } from '../../services/userApi'
import type { User } from '../../types/user'
import { formatCurrency } from '../../utils/currency'
import { getErrorMessage } from '../../utils/error'

const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

interface EditFormErrors {
  name?: string
  email?: string
  password?: string
  confirmPassword?: string
}

export default function UserProfilePage() {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()
  const authUser = useAppSelector(selectCurrentUser)

  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const [isEditOpen, setIsEditOpen] = useState(false)
  const [editName, setEditName] = useState('')
  const [editEmail, setEditEmail] = useState('')
  const [editPassword, setEditPassword] = useState('')
  const [editConfirmPassword, setEditConfirmPassword] = useState('')
  const [editFieldErrors, setEditFieldErrors] = useState<EditFormErrors>({})
  const [editError, setEditError] = useState<string | null>(null)
  const [editSuccessMessage, setEditSuccessMessage] = useState<string | null>(null)
  const [isEditSubmitting, setIsEditSubmitting] = useState(false)

  const [isDeleteOpen, setIsDeleteOpen] = useState(false)
  const [deleteError, setDeleteError] = useState<string | null>(null)
  const [isDeleteSubmitting, setIsDeleteSubmitting] = useState(false)

  const loadUser = useCallback(async () => {
    if (!authUser) {
      setUser(null)
      setError('Authenticated user information is unavailable.')
      return
    }

    setIsLoading(true)
    setError(null)

    try {
      const response = await getUserById(authUser.user_id)
      setUser(response)
    } catch (err) {
      setUser(null)
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [authUser])

  useEffect(() => {
    void loadUser()
  }, [loadUser])

  function openEditModal() {
    if (!user) {
      return
    }

    setEditName(user.name)
    setEditEmail(user.email)
    setEditPassword('')
    setEditConfirmPassword('')
    setEditFieldErrors({})
    setEditError(null)
    setEditSuccessMessage(null)
    setIsEditOpen(true)
  }

  function closeEditModal() {
    if (isEditSubmitting) {
      return
    }

    setIsEditOpen(false)
  }

  function validateEditForm(): EditFormErrors {
    const errors: EditFormErrors = {}
    const trimmedName = editName.trim()
    const trimmedEmail = editEmail.trim()

    if (!trimmedName) {
      errors.name = 'Name is required.'
    }

    if (!trimmedEmail) {
      errors.email = 'Email is required.'
    } else if (!EMAIL_PATTERN.test(trimmedEmail)) {
      errors.email = 'Enter a valid email address.'
    }

    if (!editPassword) {
      errors.password = 'Password is required.'
    }

    if (!editConfirmPassword) {
      errors.confirmPassword = 'Please confirm your password.'
    } else if (editPassword !== editConfirmPassword) {
      errors.confirmPassword = 'Passwords do not match.'
    }

    return errors
  }

  async function handleEditSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()

    if (!authUser || isEditSubmitting) {
      return
    }

    setEditError(null)
    setEditSuccessMessage(null)

    const errors = validateEditForm()
    setEditFieldErrors(errors)

    if (Object.keys(errors).length > 0) {
      return
    }

    setIsEditSubmitting(true)

    try {
      const response = await updateUser(authUser.user_id, {
        name: editName.trim(),
        email: editEmail.trim(),
        password: editPassword,
      })

      setEditSuccessMessage(response.message)
      await loadUser()
    } catch (err) {
      setEditError(getErrorMessage(err))
    } finally {
      setIsEditSubmitting(false)
    }
  }

  function openDeleteModal() {
    setDeleteError(null)
    setIsDeleteOpen(true)
  }

  function closeDeleteModal() {
    if (isDeleteSubmitting) {
      return
    }

    setIsDeleteOpen(false)
  }

  async function handleDeleteConfirm() {
    if (!authUser || isDeleteSubmitting) {
      return
    }

    setDeleteError(null)
    setIsDeleteSubmitting(true)

    try {
      await deleteUser(authUser.user_id)
      dispatch(logout())
      navigate('/login', { replace: true })
    } catch (err) {
      setDeleteError(getErrorMessage(err))
    } finally {
      setIsDeleteSubmitting(false)
    }
  }

  return (
    <PageContainer
      title="User Profile"
      description="View and manage your PayLater account."
    >
      {isLoading ? <Loader label="Loading profile" /> : null}

      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load profile" message={error} />
          <Button type="button" variant="secondary" onClick={() => void loadUser()}>
            Retry
          </Button>
        </div>
      ) : null}

      {!isLoading && !error && user ? (
        <div className="pl-user-profile">
          <header className="pl-user-profile__header">
            <h2 className="pl-user-profile__name">{user.name}</h2>
            <p className="pl-user-profile__email">{user.email}</p>
          </header>

          <div className="pl-user-profile__actions">
            <Button type="button" variant="secondary" onClick={openEditModal}>
              Edit Profile
            </Button>
            <Button type="button" variant="danger" onClick={openDeleteModal}>
              Delete Account
            </Button>
          </div>

          <div className="pl-user-profile__sections">
            <Card title="Personal Information">
              <dl className="pl-profile-fields">
                <ProfileField label="Full Name" value={user.name} />
                <ProfileField label="Email" value={user.email} />
                <ProfileField label="User ID" value={String(user.id)} />
              </dl>
            </Card>

            <Card title="Account Information">
              <dl className="pl-profile-fields">
                <ProfileField
                  label="Credit Limit"
                  value={formatCurrency(user.credit_limit)}
                />
                <ProfileField
                  label="Current Due"
                  value={formatCurrency(user.current_due)}
                />
              </dl>
            </Card>
          </div>
        </div>
      ) : null}

      <Modal
        open={isEditOpen}
        title="Edit Profile"
        onClose={closeEditModal}
        footer={
          <>
            <Button
              type="button"
              variant="secondary"
              onClick={closeEditModal}
              disabled={isEditSubmitting}
            >
              Cancel
            </Button>
            <Button
              type="submit"
              form="user-profile-edit-form"
              loading={isEditSubmitting}
            >
              Save Changes
            </Button>
          </>
        }
      >
        <form
          id="user-profile-edit-form"
          className="pl-modal__form"
          onSubmit={(event) => void handleEditSubmit(event)}
          noValidate
        >
          {editSuccessMessage ? (
            <div className="pl-user-payback__success" role="status">
              <strong className="pl-user-payback__success-title">Success</strong>
              <p className="pl-user-payback__success-message">
                {editSuccessMessage}
              </p>
            </div>
          ) : null}

          {editError ? (
            <ErrorMessage title="Unable to update profile" message={editError} />
          ) : null}

          <Input
            label="Full Name"
            id="edit-profile-name"
            name="name"
            type="text"
            autoComplete="name"
            value={editName}
            onChange={(event) => setEditName(event.target.value)}
            error={editFieldErrors.name}
            disabled={isEditSubmitting}
            required
          />

          <Input
            label="Email"
            id="edit-profile-email"
            name="email"
            type="email"
            autoComplete="email"
            value={editEmail}
            onChange={(event) => setEditEmail(event.target.value)}
            error={editFieldErrors.email}
            disabled={isEditSubmitting}
            required
          />

          <Input
            label="New Password"
            id="edit-profile-password"
            name="password"
            type="password"
            autoComplete="new-password"
            value={editPassword}
            onChange={(event) => setEditPassword(event.target.value)}
            error={editFieldErrors.password}
            disabled={isEditSubmitting}
            required
          />

          <Input
            label="Confirm Password"
            id="edit-profile-confirm-password"
            name="confirmPassword"
            type="password"
            autoComplete="new-password"
            value={editConfirmPassword}
            onChange={(event) => setEditConfirmPassword(event.target.value)}
            error={editFieldErrors.confirmPassword}
            disabled={isEditSubmitting}
            required
          />
        </form>
      </Modal>

      <Modal
        open={isDeleteOpen}
        title="Delete Account"
        onClose={closeDeleteModal}
        footer={
          <>
            <Button
              type="button"
              variant="secondary"
              onClick={closeDeleteModal}
              disabled={isDeleteSubmitting}
            >
              Cancel
            </Button>
            <Button
              type="button"
              variant="danger"
              loading={isDeleteSubmitting}
              onClick={() => void handleDeleteConfirm()}
            >
              Delete Account
            </Button>
          </>
        }
      >
        {deleteError ? (
          <ErrorMessage title="Unable to delete account" message={deleteError} />
        ) : null}

        <p className="pl-modal__confirm-text">
          Are you sure you want to permanently delete your account
          {user ? ` (${user.email})` : ''}? This action cannot be undone.
        </p>
      </Modal>
    </PageContainer>
  )
}

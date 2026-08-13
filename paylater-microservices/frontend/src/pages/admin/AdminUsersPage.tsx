import { useCallback, useEffect, useMemo, useState, type FormEvent } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import StatCard from '../../components/ui/StatCard'
import Modal from '../../components/ui/Modal'
import Table, { type TableColumn } from '../../components/ui/Table'
import { deleteUser, getUsers, updateUser } from '../../services/userApi'
import { sumUserTotalDue } from '../../features/admin/adminUserUtils'
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

export default function AdminUsersPage() {
  const [users, setUsers] = useState<User[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const [editingUser, setEditingUser] = useState<User | null>(null)
  const [editName, setEditName] = useState('')
  const [editEmail, setEditEmail] = useState('')
  const [editPassword, setEditPassword] = useState('')
  const [editConfirmPassword, setEditConfirmPassword] = useState('')
  const [editFieldErrors, setEditFieldErrors] = useState<EditFormErrors>({})
  const [editError, setEditError] = useState<string | null>(null)
  const [editSuccessMessage, setEditSuccessMessage] = useState<string | null>(null)
  const [isEditSubmitting, setIsEditSubmitting] = useState(false)

  const [deletingUser, setDeletingUser] = useState<User | null>(null)
  const [deleteError, setDeleteError] = useState<string | null>(null)
  const [isDeleteSubmitting, setIsDeleteSubmitting] = useState(false)

  const loadUsers = useCallback(async () => {
    setIsLoading(true)
    setError(null)

    try {
      const response = await getUsers()
      setUsers(response ?? [])
    } catch (err) {
      setUsers([])
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [])

  useEffect(() => {
    void loadUsers()
  }, [loadUsers])

  function openEditModal(user: User) {
    setEditingUser(user)
    setEditName(user.name)
    setEditEmail(user.email)
    setEditPassword('')
    setEditConfirmPassword('')
    setEditFieldErrors({})
    setEditError(null)
    setEditSuccessMessage(null)
  }

  function closeEditModal() {
    if (isEditSubmitting) {
      return
    }

    setEditingUser(null)
  }

  function openDeleteModal(user: User) {
    setDeletingUser(user)
    setDeleteError(null)
  }

  function closeDeleteModal() {
    if (isDeleteSubmitting) {
      return
    }

    setDeletingUser(null)
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
      errors.confirmPassword = 'Please confirm the password.'
    } else if (editPassword !== editConfirmPassword) {
      errors.confirmPassword = 'Passwords do not match.'
    }

    return errors
  }

  async function handleEditSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()

    if (!editingUser || isEditSubmitting) {
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
      const response = await updateUser(editingUser.id, {
        name: editName.trim(),
        email: editEmail.trim(),
        password: editPassword,
      })

      setEditSuccessMessage(response.message)
      await loadUsers()
    } catch (err) {
      setEditError(getErrorMessage(err))
    } finally {
      setIsEditSubmitting(false)
    }
  }

  async function handleDeleteConfirm() {
    if (!deletingUser || isDeleteSubmitting) {
      return
    }

    setDeleteError(null)
    setIsDeleteSubmitting(true)

    try {
      await deleteUser(deletingUser.id)
      setDeletingUser(null)
      await loadUsers()
    } catch (err) {
      setDeleteError(getErrorMessage(err))
    } finally {
      setIsDeleteSubmitting(false)
    }
  }

  const userColumns = useMemo<TableColumn<User>[]>(
    () => [
      { key: 'id', header: 'ID' },
      { key: 'name', header: 'Name' },
      { key: 'email', header: 'Email' },
      {
        key: 'credit_limit',
        header: 'Credit Limit',
        render: (row) => formatCurrency(row.credit_limit),
      },
      {
        key: 'current_due',
        header: 'Current Due',
        render: (row) => formatCurrency(row.current_due),
      },
      {
        key: 'actions',
        header: 'Actions',
        render: (row) => (
          <div className="pl-table-actions">
            <Button
              type="button"
              variant="secondary"
              size="small"
              onClick={() => openEditModal(row)}
            >
              Edit
            </Button>
            <Button
              type="button"
              variant="danger"
              size="small"
              onClick={() => openDeleteModal(row)}
            >
              Delete
            </Button>
          </div>
        ),
      },
    ],
    [],
  )

  return (
    <PageContainer
      title="Users"
      description="Manage PayLater user accounts."
    >
      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load users" message={error} />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadUsers()}
          >
            Retry
          </Button>
        </div>
      ) : (
        <>
          {!isLoading ? (
            <section
              className="pl-admin-stats pl-admin-users-stats"
              aria-label="User summary statistics"
            >
              <StatCard label="Total Users" value={String(users.length)} />
              <StatCard
                label="Total Due Amount"
                value={formatCurrency(sumUserTotalDue(users).toFixed(2))}
              />
            </section>
          ) : null}

          <Card title="All Users">
          <Table
            columns={userColumns}
            data={users}
            loading={isLoading}
            emptyMessage="No users found."
          />
        </Card>
        </>
      )}

      <Modal
        open={editingUser !== null}
        title={editingUser ? `Edit User #${editingUser.id}` : 'Edit User'}
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
              form="admin-edit-user-form"
              loading={isEditSubmitting}
            >
              Save Changes
            </Button>
          </>
        }
      >
        <form
          id="admin-edit-user-form"
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
            <ErrorMessage title="Unable to update user" message={editError} />
          ) : null}

          <Input
            label="Name"
            id="admin-edit-user-name"
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
            id="admin-edit-user-email"
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
            id="admin-edit-user-password"
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
            id="admin-edit-user-confirm-password"
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
        open={deletingUser !== null}
        title={deletingUser ? `Delete User #${deletingUser.id}` : 'Delete User'}
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
              Delete User
            </Button>
          </>
        }
      >
        {deleteError ? (
          <ErrorMessage title="Unable to delete user" message={deleteError} />
        ) : null}

        <p className="pl-modal__confirm-text">
          Are you sure you want to permanently delete{' '}
          {deletingUser ? `${deletingUser.name} (${deletingUser.email})` : 'this user'}
          ? This action cannot be undone.
        </p>
      </Modal>
    </PageContainer>
  )
}

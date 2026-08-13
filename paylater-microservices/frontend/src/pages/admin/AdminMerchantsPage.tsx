import { useCallback, useEffect, useMemo, useState, type FormEvent } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import StatCard from '../../components/ui/StatCard'
import Modal from '../../components/ui/Modal'
import Table, { type TableColumn } from '../../components/ui/Table'
import {
  deleteMerchant,
  getMerchants,
  updateCommission,
  updateMerchant,
} from '../../services/merchantApi'
import type { Merchant } from '../../types/merchant'
import { formatCommission } from '../../utils/commission'
import { getErrorMessage } from '../../utils/error'

const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

interface EditFormErrors {
  merchantName?: string
  email?: string
  phone?: string
  commission?: string
  password?: string
  confirmPassword?: string
}

interface CommissionFormErrors {
  commission?: string
}

export default function AdminMerchantsPage() {
  const [merchants, setMerchants] = useState<Merchant[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const [editingMerchant, setEditingMerchant] = useState<Merchant | null>(null)
  const [editMerchantName, setEditMerchantName] = useState('')
  const [editEmail, setEditEmail] = useState('')
  const [editPhone, setEditPhone] = useState('')
  const [editCommission, setEditCommission] = useState('')
  const [editPassword, setEditPassword] = useState('')
  const [editConfirmPassword, setEditConfirmPassword] = useState('')
  const [editFieldErrors, setEditFieldErrors] = useState<EditFormErrors>({})
  const [editError, setEditError] = useState<string | null>(null)
  const [editSuccessMessage, setEditSuccessMessage] = useState<string | null>(null)
  const [isEditSubmitting, setIsEditSubmitting] = useState(false)

  const [commissionMerchant, setCommissionMerchant] = useState<Merchant | null>(null)
  const [commissionValue, setCommissionValue] = useState('')
  const [commissionFieldErrors, setCommissionFieldErrors] =
    useState<CommissionFormErrors>({})
  const [commissionError, setCommissionError] = useState<string | null>(null)
  const [commissionSuccessMessage, setCommissionSuccessMessage] = useState<
    string | null
  >(null)
  const [isCommissionSubmitting, setIsCommissionSubmitting] = useState(false)

  const [deletingMerchant, setDeletingMerchant] = useState<Merchant | null>(null)
  const [deleteError, setDeleteError] = useState<string | null>(null)
  const [isDeleteSubmitting, setIsDeleteSubmitting] = useState(false)

  const loadMerchants = useCallback(async () => {
    setIsLoading(true)
    setError(null)

    try {
      const response = await getMerchants()
      setMerchants(response)
    } catch (err) {
      setMerchants([])
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [])

  useEffect(() => {
    void loadMerchants()
  }, [loadMerchants])

  function openEditModal(merchant: Merchant) {
    setEditingMerchant(merchant)
    setEditMerchantName(merchant.merchant_name)
    setEditEmail(merchant.email)
    setEditPhone(merchant.phone)
    setEditCommission(merchant.commission)
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

    setEditingMerchant(null)
  }

  function openCommissionModal(merchant: Merchant) {
    setCommissionMerchant(merchant)
    setCommissionValue(merchant.commission)
    setCommissionFieldErrors({})
    setCommissionError(null)
    setCommissionSuccessMessage(null)
  }

  function closeCommissionModal() {
    if (isCommissionSubmitting) {
      return
    }

    setCommissionMerchant(null)
  }

  function openDeleteModal(merchant: Merchant) {
    setDeletingMerchant(merchant)
    setDeleteError(null)
  }

  function closeDeleteModal() {
    if (isDeleteSubmitting) {
      return
    }

    setDeletingMerchant(null)
  }

  function validateEditForm(): EditFormErrors {
    const errors: EditFormErrors = {}
    const trimmedMerchantName = editMerchantName.trim()
    const trimmedEmail = editEmail.trim()
    const trimmedPhone = editPhone.trim()

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

    if (!editCommission.trim()) {
      errors.commission = 'Commission is required.'
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

    if (!editingMerchant || isEditSubmitting) {
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
      const response = await updateMerchant(editingMerchant.id, {
        merchant_name: editMerchantName.trim(),
        phone: editPhone.trim(),
        commission: editCommission.trim(),
        email: editEmail.trim(),
        password: editPassword,
      })

      setEditSuccessMessage(response.message)
      await loadMerchants()
    } catch (err) {
      setEditError(getErrorMessage(err))
    } finally {
      setIsEditSubmitting(false)
    }
  }

  async function handleCommissionSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()

    if (!commissionMerchant || isCommissionSubmitting) {
      return
    }

    setCommissionError(null)
    setCommissionSuccessMessage(null)

    const errors: CommissionFormErrors = {}

    if (!commissionValue.trim()) {
      errors.commission = 'Commission is required.'
    }

    setCommissionFieldErrors(errors)

    if (Object.keys(errors).length > 0) {
      return
    }

    setIsCommissionSubmitting(true)

    try {
      const response = await updateCommission(commissionMerchant.id, {
        commission: commissionValue.trim(),
      })

      setCommissionSuccessMessage(response.message)
      await loadMerchants()
    } catch (err) {
      setCommissionError(getErrorMessage(err))
    } finally {
      setIsCommissionSubmitting(false)
    }
  }

  async function handleDeleteConfirm() {
    if (!deletingMerchant || isDeleteSubmitting) {
      return
    }

    setDeleteError(null)
    setIsDeleteSubmitting(true)

    try {
      await deleteMerchant(deletingMerchant.id)
      setDeletingMerchant(null)
      await loadMerchants()
    } catch (err) {
      setDeleteError(getErrorMessage(err))
    } finally {
      setIsDeleteSubmitting(false)
    }
  }

  const merchantColumns = useMemo<TableColumn<Merchant>[]>(
    () => [
      { key: 'id', header: 'ID' },
      { key: 'merchant_name', header: 'Merchant Name' },
      { key: 'email', header: 'Email' },
      { key: 'phone', header: 'Phone' },
      {
        key: 'commission',
        header: 'Commission',
        render: (row) => formatCommission(row.commission),
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
              variant="outline"
              size="small"
              onClick={() => openCommissionModal(row)}
            >
              Commission
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
      title="Merchants"
      description="Manage PayLater merchant accounts."
    >
      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load merchants" message={error} />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadMerchants()}
          >
            Retry
          </Button>
        </div>
      ) : (
        <>
          {!isLoading ? (
            <section
              className="pl-admin-stats pl-admin-merchants-stats"
              aria-label="Merchant summary statistics"
            >
              <StatCard
                label="Total Merchants"
                value={String(merchants.length)}
              />
            </section>
          ) : null}

          <Card title="All Merchants">
            <Table
              columns={merchantColumns}
              data={merchants}
              loading={isLoading}
              emptyMessage="No merchants found."
            />
          </Card>
        </>
      )}

      <Modal
        open={editingMerchant !== null}
        title={
          editingMerchant
            ? `Edit Merchant #${editingMerchant.id}`
            : 'Edit Merchant'
        }
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
              form="admin-edit-merchant-form"
              loading={isEditSubmitting}
            >
              Save Changes
            </Button>
          </>
        }
      >
        <form
          id="admin-edit-merchant-form"
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
            <ErrorMessage title="Unable to update merchant" message={editError} />
          ) : null}

          <Input
            label="Merchant Name"
            id="admin-edit-merchant-name"
            name="merchantName"
            type="text"
            autoComplete="organization"
            value={editMerchantName}
            onChange={(event) => setEditMerchantName(event.target.value)}
            error={editFieldErrors.merchantName}
            disabled={isEditSubmitting}
            required
          />

          <Input
            label="Email"
            id="admin-edit-merchant-email"
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
            label="Phone"
            id="admin-edit-merchant-phone"
            name="phone"
            type="tel"
            autoComplete="tel"
            value={editPhone}
            onChange={(event) => setEditPhone(event.target.value)}
            error={editFieldErrors.phone}
            disabled={isEditSubmitting}
            required
          />

          <Input
            label="Commission (%)"
            id="admin-edit-merchant-commission"
            name="commission"
            type="number"
            inputMode="decimal"
            step="0.01"
            value={editCommission}
            onChange={(event) => setEditCommission(event.target.value)}
            error={editFieldErrors.commission}
            disabled={isEditSubmitting}
            required
          />

          <Input
            label="New Password"
            id="admin-edit-merchant-password"
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
            id="admin-edit-merchant-confirm-password"
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
        open={commissionMerchant !== null}
        title={
          commissionMerchant
            ? `Update Commission — ${commissionMerchant.merchant_name}`
            : 'Update Commission'
        }
        onClose={closeCommissionModal}
        footer={
          <>
            <Button
              type="button"
              variant="secondary"
              onClick={closeCommissionModal}
              disabled={isCommissionSubmitting}
            >
              Cancel
            </Button>
            <Button
              type="submit"
              form="admin-update-commission-form"
              loading={isCommissionSubmitting}
            >
              Update Commission
            </Button>
          </>
        }
      >
        <form
          id="admin-update-commission-form"
          className="pl-modal__form"
          onSubmit={(event) => void handleCommissionSubmit(event)}
          noValidate
        >
          {commissionSuccessMessage ? (
            <div className="pl-user-payback__success" role="status">
              <strong className="pl-user-payback__success-title">Success</strong>
              <p className="pl-user-payback__success-message">
                {commissionSuccessMessage}
              </p>
            </div>
          ) : null}

          {commissionError ? (
            <ErrorMessage
              title="Unable to update commission"
              message={commissionError}
            />
          ) : null}

          <Input
            label="Commission (%)"
            id="admin-update-commission"
            name="commission"
            type="number"
            inputMode="decimal"
            step="0.01"
            value={commissionValue}
            onChange={(event) => setCommissionValue(event.target.value)}
            error={commissionFieldErrors.commission}
            disabled={isCommissionSubmitting}
            required
          />
        </form>
      </Modal>

      <Modal
        open={deletingMerchant !== null}
        title={
          deletingMerchant
            ? `Delete Merchant #${deletingMerchant.id}`
            : 'Delete Merchant'
        }
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
              Delete Merchant
            </Button>
          </>
        }
      >
        {deleteError ? (
          <ErrorMessage title="Unable to delete merchant" message={deleteError} />
        ) : null}

        <p className="pl-modal__confirm-text">
          Are you sure you want to permanently delete{' '}
          {deletingMerchant
            ? `${deletingMerchant.merchant_name} (${deletingMerchant.email})`
            : 'this merchant'}
          ? This action cannot be undone.
        </p>
      </Modal>
    </PageContainer>
  )
}

import { useCallback, useEffect, useState, type FormEvent } from 'react'

import { useAppSelector } from '../../app/hooks'
import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Input from '../../components/common/Input'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import Table, { type TableColumn } from '../../components/ui/Table'
import { selectCurrentUser } from '../../features/auth/authSelectors'
import { getMerchants } from '../../services/merchantApi'
import {
  createTransaction,
  getTransactionsByUser,
} from '../../services/transactionApi'
import type { Merchant } from '../../types/merchant'
import type {
  CreateTransactionResponse,
  Transaction,
} from '../../types/transaction'
import { formatCurrency } from '../../utils/currency'
import { formatDateTime } from '../../utils/date'
import { getErrorMessage } from '../../utils/error'

const transactionColumns: TableColumn<Transaction>[] = [
  { key: 'id', header: 'Transaction ID' },
  {
    key: 'amount',
    header: 'Amount',
    render: (row) => formatCurrency(row.amount),
  },
  { key: 'transaction_type', header: 'Type' },
  {
    key: 'merchant_id',
    header: 'Merchant',
    render: (row) =>
      row.merchant_id !== undefined ? `Merchant #${row.merchant_id}` : '—',
  },
  {
    key: 'transaction_date',
    header: 'Date',
    render: (row) =>
      row.transaction_date ? formatDateTime(row.transaction_date) : '—',
  },
]

interface FormErrors {
  merchantId?: string
  amount?: string
}

function validateForm(merchantId: string, amount: string): FormErrors {
  const errors: FormErrors = {}
  const trimmedAmount = amount.trim()

  if (!merchantId) {
    errors.merchantId = 'Please select a merchant.'
  }

  if (!trimmedAmount) {
    errors.amount = 'Amount is required.'
  } else {
    const parsedAmount = Number(trimmedAmount)
    if (!Number.isFinite(parsedAmount) || parsedAmount <= 0) {
      errors.amount = 'Amount must be greater than 0.'
    }
  }

  return errors
}

export default function UserTransactionsPage() {
  const authUser = useAppSelector(selectCurrentUser)
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [merchants, setMerchants] = useState<Merchant[]>([])
  const [isLoadingTransactions, setIsLoadingTransactions] = useState(false)
  const [isLoadingMerchants, setIsLoadingMerchants] = useState(false)
  const [transactionsError, setTransactionsError] = useState<string | null>(null)
  const [merchantsError, setMerchantsError] = useState<string | null>(null)

  const [merchantId, setMerchantId] = useState('')
  const [amount, setAmount] = useState('')
  const [fieldErrors, setFieldErrors] = useState<FormErrors>({})
  const [submitError, setSubmitError] = useState<string | null>(null)
  const [successResult, setSuccessResult] =
    useState<CreateTransactionResponse | null>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)

  const loadTransactions = useCallback(async () => {
    if (!authUser) {
      setTransactions([])
      setTransactionsError('Authenticated user information is unavailable.')
      return
    }

    setIsLoadingTransactions(true)
    setTransactionsError(null)

    try {
      const response = await getTransactionsByUser(authUser.user_id)
      setTransactions(response)
    } catch (err) {
      setTransactions([])
      setTransactionsError(getErrorMessage(err))
    } finally {
      setIsLoadingTransactions(false)
    }
  }, [authUser])

  const loadMerchants = useCallback(async () => {
    setIsLoadingMerchants(true)
    setMerchantsError(null)

    try {
      const response = await getMerchants()
      setMerchants(response)
    } catch (err) {
      setMerchants([])
      setMerchantsError(getErrorMessage(err))
    } finally {
      setIsLoadingMerchants(false)
    }
  }, [])

  useEffect(() => {
    void loadTransactions()
    void loadMerchants()
  }, [loadTransactions, loadMerchants])

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()

    if (!authUser || isSubmitting) {
      return
    }

    setSubmitError(null)
    setSuccessResult(null)

    const errors = validateForm(merchantId, amount)
    setFieldErrors(errors)

    if (Object.keys(errors).length > 0) {
      return
    }

    setIsSubmitting(true)
    try {
      const response = await createTransaction({
        merchant_id: Number(merchantId),
        amount: amount.trim(),
      })

      setSuccessResult(response)
      setMerchantId('')
      setAmount('')
      setFieldErrors({})
      await loadTransactions()
    } catch (err) {
      setSubmitError(getErrorMessage(err))
    } finally {
      setIsSubmitting(false)
    }
  }

  const isFormDisabled =
    isSubmitting || isLoadingMerchants || merchants.length === 0

  return (
    <PageContainer
      title="Transactions"
      description="Your PayLater transaction history."
    >
      <div className="pl-user-transactions">
        <Card title="New Purchase">
          {isLoadingMerchants ? (
            <Loader label="Loading merchants" />
          ) : null}

          {!isLoadingMerchants && merchantsError ? (
            <div className="pl-dashboard-error">
              <ErrorMessage
                title="Unable to load merchants"
                message={merchantsError}
              />
              <Button
                type="button"
                variant="secondary"
                onClick={() => void loadMerchants()}
              >
                Retry
              </Button>
            </div>
          ) : null}

          {!isLoadingMerchants && !merchantsError && merchants.length === 0 ? (
            <p className="pl-input__helper">No merchants are available.</p>
          ) : null}

          {!isLoadingMerchants && !merchantsError && merchants.length > 0 ? (
            <form
              className="pl-user-transactions__form"
              onSubmit={(event) => void handleSubmit(event)}
              noValidate
            >
              {successResult ? (
                <div className="pl-user-payback__success" role="status">
                  <strong className="pl-user-payback__success-title">
                    Success
                  </strong>
                  <p className="pl-user-payback__success-message">
                    {successResult.message}
                  </p>
                  <p className="pl-user-payback__success-message">
                    Amount: {formatCurrency(successResult.transaction_amount)} ·
                    Updated current due:{' '}
                    {formatCurrency(successResult.updated_current_due)} ·
                    Available credit:{' '}
                    {formatCurrency(successResult.available_credit)}
                  </p>
                </div>
              ) : null}

              {submitError ? (
                <ErrorMessage
                  title="Transaction failed"
                  message={submitError}
                />
              ) : null}

              <div className="pl-input">
                <label className="pl-input__label" htmlFor="transaction-merchant">
                  Merchant<span aria-hidden="true"> *</span>
                </label>
                <select
                  id="transaction-merchant"
                  name="merchantId"
                  className={`pl-input__field${fieldErrors.merchantId ? ' pl-input__field--error' : ''}`}
                  value={merchantId}
                  onChange={(event) => setMerchantId(event.target.value)}
                  disabled={isFormDisabled}
                  required
                  aria-invalid={fieldErrors.merchantId ? true : undefined}
                >
                  <option value="">Select a merchant</option>
                  {merchants.map((merchant) => (
                    <option key={merchant.id} value={String(merchant.id)}>
                      {merchant.merchant_name}
                    </option>
                  ))}
                </select>
                {fieldErrors.merchantId ? (
                  <p className="pl-input__error" role="alert">
                    {fieldErrors.merchantId}
                  </p>
                ) : null}
              </div>

              <Input
                label="Amount"
                id="transaction-amount"
                name="amount"
                type="number"
                inputMode="decimal"
                step="0.01"
                min="0"
                value={amount}
                onChange={(event) => setAmount(event.target.value)}
                error={fieldErrors.amount}
                disabled={isFormDisabled}
                required
              />

              <Button
                type="submit"
                loading={isSubmitting}
                disabled={isFormDisabled}
              >
                Create Transaction
              </Button>
            </form>
          ) : null}
        </Card>

        {!isLoadingTransactions && transactionsError ? (
          <div className="pl-dashboard-error">
            <ErrorMessage
              title="Unable to load transactions"
              message={transactionsError}
            />
            <Button
              type="button"
              variant="secondary"
              onClick={() => void loadTransactions()}
            >
              Retry
            </Button>
          </div>
        ) : (
          <Table
            columns={transactionColumns}
            data={transactions}
            loading={isLoadingTransactions}
            emptyMessage="No transactions found."
          />
        )}
      </div>
    </PageContainer>
  )
}

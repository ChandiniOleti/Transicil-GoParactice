import { useCallback, useEffect, useState } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import Table, { type TableColumn } from '../../components/ui/Table'
import { getAdmins } from '../../services/authApi'
import type { Admin } from '../../types/auth'
import { getErrorMessage } from '../../utils/error'

interface AdminRow extends Admin {
  role: 'ADMIN'
}

const adminColumns: TableColumn<AdminRow>[] = [
  { key: 'name', header: 'Name' },
  { key: 'email', header: 'Email' },
  {
    key: 'role',
    header: 'Role',
    render: (row) => row.role,
  },
]

export default function AdminAdminsPage() {
  const [admins, setAdmins] = useState<AdminRow[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadAdmins = useCallback(async () => {
    setIsLoading(true)
    setError(null)

    try {
      const response = await getAdmins()
      const rows = (response ?? []).map((admin) => ({
        ...admin,
        role: 'ADMIN' as const,
      }))
      setAdmins(rows)
    } catch (err) {
      setAdmins([])
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [])

  useEffect(() => {
    void loadAdmins()
  }, [loadAdmins])

  return (
    <PageContainer
      title="All Admins"
      description="View PayLater platform administrators."
    >
      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load admins" message={error} />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadAdmins()}
          >
            Retry
          </Button>
        </div>
      ) : (
        <Card title="All Admins">
          <Table
            columns={adminColumns}
            data={admins}
            loading={isLoading}
            emptyMessage="No admins found."
          />
        </Card>
      )}
    </PageContainer>
  )
}

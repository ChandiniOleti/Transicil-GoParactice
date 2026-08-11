import { useCallback, useEffect, useState } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import Table, { type TableColumn } from '../../components/ui/Table'
import AdminReportTables from '../../features/admin/AdminReportTables'
import AdminSummaryStats from '../../features/admin/AdminSummaryStats'
import {
  buildMerchantFeeRows,
  buildUserDueDetails,
  sumCommissionCollected,
} from '../../features/admin/adminReportUtils'
import { getMerchants } from '../../services/merchantApi'
import {
  getCreditLimitUsersReport,
  getTotalDuesReport,
  getUsersWithDueReport,
} from '../../services/reportApi'
import { getUsers } from '../../services/userApi'
import type { MerchantFeeRow } from '../../features/admin/AdminReportTables'
import type {
  CreditLimitUsersResponse,
  TotalDuesResponse,
  UserReport,
  UsersWithDueResponse,
} from '../../types/report'
import { formatCurrency } from '../../utils/currency'
import { getErrorMessage } from '../../utils/error'

interface AdminReportsData {
  totalDues: TotalDuesResponse
  usersWithDue: UsersWithDueResponse
  creditLimitUsers: CreditLimitUsersResponse
  userDueDetails: UserReport[]
  merchantFees: MerchantFeeRow[]
}

const userDueColumns: TableColumn<UserReport>[] = [
  { key: 'id', header: 'User ID' },
  { key: 'name', header: 'Name' },
  { key: 'email', header: 'Email' },
  {
    key: 'current_due',
    header: 'Current Due',
    render: (row) => formatCurrency(row.current_due),
  },
  {
    key: 'credit_limit',
    header: 'Credit Limit',
    render: (row) => formatCurrency(row.credit_limit),
  },
]

export default function AdminReportsPage() {
  const [reportsData, setReportsData] = useState<AdminReportsData | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadReports = useCallback(async () => {
    setIsLoading(true)
    setError(null)

    try {
      const [totalDues, usersWithDue, creditLimitUsers, merchants, users] =
        await Promise.all([
          getTotalDuesReport(),
          getUsersWithDueReport(),
          getCreditLimitUsersReport(),
          getMerchants(),
          getUsers(),
        ])

      const userIds = (users ?? []).map((user) => user.id)

      const [merchantFees, userDueDetails] = await Promise.all([
        buildMerchantFeeRows(merchants),
        buildUserDueDetails(userIds),
      ])

      setReportsData({
        totalDues,
        usersWithDue,
        creditLimitUsers,
        userDueDetails,
        merchantFees,
      })
    } catch (err) {
      setReportsData(null)
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [])

  useEffect(() => {
    void loadReports()
  }, [loadReports])

  const totalCommissionCollected = reportsData
    ? sumCommissionCollected(
        reportsData.merchantFees.map((row) => ({
          total_fee_collected: row.totalFeeCollected,
        })),
      )
    : '0.00'

  return (
    <PageContainer
      title="Reports"
      description="PayLater platform reports and analytics."
    >
      {isLoading ? <Loader label="Loading reports" /> : null}

      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load reports" message={error} />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadReports()}
          >
            Retry
          </Button>
        </div>
      ) : null}

      {!isLoading && !error && reportsData ? (
        <div className="pl-admin-dashboard">
          <AdminSummaryStats
            totalDues={reportsData.totalDues.total_due}
            usersWithDueCount={reportsData.usersWithDue.length}
            creditLimitUsersCount={reportsData.creditLimitUsers.length}
            merchantCount={reportsData.merchantFees.length}
            totalCommissionCollected={totalCommissionCollected}
          />

          <div className="pl-admin-report-tables">
            <Card title="Total User Dues">
              <p className="pl-admin-report-summary">
                {formatCurrency(reportsData.totalDues.total_due)}
              </p>
            </Card>

            <AdminReportTables
              usersWithDue={reportsData.usersWithDue}
              creditLimitUsers={reportsData.creditLimitUsers}
              merchantFees={reportsData.merchantFees}
            />

            <Card title="User Due Details">
              <Table
                columns={userDueColumns}
                data={reportsData.userDueDetails}
                emptyMessage="No user due details available."
              />
            </Card>
          </div>
        </div>
      ) : null}
    </PageContainer>
  )
}

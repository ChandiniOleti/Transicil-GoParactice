import { useCallback, useEffect, useState } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import Card from '../../components/ui/Card'
import Table, { type TableColumn } from '../../components/ui/Table'
import AdminReportTables from '../../features/admin/AdminReportTables'
import AdminSummaryStats, {
  type AdminSummaryStatKey,
} from '../../features/admin/AdminSummaryStats'
import CustomerDueSearch from '../../features/admin/CustomerDueSearch'
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
  users: NonNullable<Awaited<ReturnType<typeof getUsers>>>
}

const REPORT_SECTION_IDS: Record<AdminSummaryStatKey, string> = {
  totalDues: 'report-total-user-dues',
  usersWithDue: 'report-users-with-due',
  creditLimitUsers: 'report-credit-limit-users',
  merchantCount: 'report-merchant-commission',
  totalCommissionCollected: 'report-merchant-commission',
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

function scrollToReportSection(sectionId: string) {
  document.getElementById(sectionId)?.scrollIntoView({
    behavior: 'smooth',
    block: 'start',
  })
}

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
        users: users ?? [],
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

  const clickableStats = reportsData
    ? {
        totalDues: true,
        usersWithDue: reportsData.usersWithDue.length > 0,
        creditLimitUsers: reportsData.creditLimitUsers.length > 0,
        merchantCount: reportsData.merchantFees.length > 0,
        totalCommissionCollected: reportsData.merchantFees.length > 0,
      }
    : undefined

  function handleStatClick(key: AdminSummaryStatKey) {
    scrollToReportSection(REPORT_SECTION_IDS[key])
  }

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
            clickableStats={clickableStats}
            onStatClick={handleStatClick}
          />

          <div id="report-customer-due-search" className="pl-admin-report-section">
            <CustomerDueSearch users={reportsData.users} />
          </div>

          <div className="pl-admin-report-tables">
            <div id="report-total-user-dues" className="pl-admin-report-section">
              <Card title="Total User Dues">
                <p className="pl-admin-report-summary">
                  {formatCurrency(reportsData.totalDues.total_due)}
                </p>
              </Card>
            </div>

            <AdminReportTables
              usersWithDue={reportsData.usersWithDue}
              creditLimitUsers={reportsData.creditLimitUsers}
              merchantFees={reportsData.merchantFees}
            />

            <div id="report-user-due-details" className="pl-admin-report-section">
              <Card title="User Due Details">
                <Table
                  columns={userDueColumns}
                  data={reportsData.userDueDetails}
                  emptyMessage="No user due details available."
                />
              </Card>
            </div>
          </div>
        </div>
      ) : null}
    </PageContainer>
  )
}

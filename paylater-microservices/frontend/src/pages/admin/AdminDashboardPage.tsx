import { useCallback, useEffect, useState } from 'react'

import Button from '../../components/common/Button'
import ErrorMessage from '../../components/common/ErrorMessage'
import Loader from '../../components/common/Loader'
import PageContainer from '../../components/layout/PageContainer'
import AdminDashboardOverview from '../../features/admin/AdminDashboardOverview'
import AdminReportTables from '../../features/admin/AdminReportTables'
import {
  buildMerchantFeeRows,
} from '../../features/admin/adminReportUtils'
import { getMerchants } from '../../services/merchantApi'
import {
  getCreditLimitUsersReport,
  getTotalDuesReport,
  getUsersWithDueReport,
} from '../../services/reportApi'
import { getUsers } from '../../services/userApi'
import type {
  CreditLimitUsersResponse,
  TotalDuesResponse,
  UsersWithDueResponse,
} from '../../types/report'
import type { MerchantFeeRow } from '../../features/admin/AdminReportTables'
import { getErrorMessage } from '../../utils/error'

interface AdminDashboardData {
  totalUsers: number
  totalMerchants: number
  totalDues: TotalDuesResponse
  usersWithDue: UsersWithDueResponse
  creditLimitUsers: CreditLimitUsersResponse
  merchantFees: MerchantFeeRow[]
}

export default function AdminDashboardPage() {
  const [dashboardData, setDashboardData] = useState<AdminDashboardData | null>(
    null,
  )
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadDashboard = useCallback(async () => {
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

      const merchantFees = await buildMerchantFeeRows(merchants)

      setDashboardData({
        totalUsers: (users ?? []).length,
        totalMerchants: merchants.length,
        totalDues,
        usersWithDue,
        creditLimitUsers,
        merchantFees,
      })
    } catch (err) {
      setDashboardData(null)
      setError(getErrorMessage(err))
    } finally {
      setIsLoading(false)
    }
  }, [])

  useEffect(() => {
    void loadDashboard()
  }, [loadDashboard])

  return (
    <PageContainer
      title="Admin Dashboard"
      description="PayLater platform overview and report summaries."
    >
      {isLoading ? <Loader label="Loading dashboard data" /> : null}

      {!isLoading && error ? (
        <div className="pl-dashboard-error">
          <ErrorMessage title="Unable to load dashboard" message={error} />
          <Button
            type="button"
            variant="secondary"
            onClick={() => void loadDashboard()}
          >
            Retry
          </Button>
        </div>
      ) : null}

      {!isLoading && !error && dashboardData ? (
        <div className="pl-admin-dashboard">
          <AdminDashboardOverview
            totalUsers={dashboardData.totalUsers}
            totalMerchants={dashboardData.totalMerchants}
            totalOutstandingDue={dashboardData.totalDues.total_due}
          />
          <AdminReportTables
            usersWithDue={dashboardData.usersWithDue}
            creditLimitUsers={dashboardData.creditLimitUsers}
            merchantFees={dashboardData.merchantFees}
          />
        </div>
      ) : null}
    </PageContainer>
  )
}

import { useMemo, useState } from 'react'

import Input from '../../components/common/Input'
import Card from '../../components/ui/Card'
import Table, { type TableColumn } from '../../components/ui/Table'
import { filterUsersByName } from './adminUserUtils'
import type { User } from '../../types/user'
import { formatCurrency } from '../../utils/currency'

export interface CustomerDueSearchProps {
  users: User[]
}

interface CustomerDueRow {
  id: number
  name: string
  currentDue: string
}

const customerDueColumns: TableColumn<CustomerDueRow>[] = [
  { key: 'name', header: 'Customer Name' },
  {
    key: 'currentDue',
    header: 'Current Due',
    render: (row) => formatCurrency(row.currentDue),
  },
]

export default function CustomerDueSearch({ users }: CustomerDueSearchProps) {
  const [searchQuery, setSearchQuery] = useState('')

  const searchResults = useMemo<CustomerDueRow[]>(() => {
    return filterUsersByName(users, searchQuery).map((user) => ({
      id: user.id,
      name: user.name,
      currentDue: user.current_due,
    }))
  }, [searchQuery, users])

  const emptyMessage = searchQuery.trim()
    ? 'No customers match your search.'
    : 'Enter a customer name to search.'

  return (
    <Card title="Customer Due Search">
      <div className="pl-customer-due-search">
        <Input
          label="Search by customer name"
          id="customer-due-search"
          name="customerDueSearch"
          type="search"
          autoComplete="off"
          value={searchQuery}
          onChange={(event) => setSearchQuery(event.target.value)}
          placeholder="Type a customer name"
        />

        <Table
          columns={customerDueColumns}
          data={searchQuery.trim() ? searchResults : []}
          emptyMessage={emptyMessage}
        />
      </div>
    </Card>
  )
}

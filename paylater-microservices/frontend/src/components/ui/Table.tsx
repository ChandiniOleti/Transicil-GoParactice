import type { ReactNode } from 'react'

import EmptyState from '../common/EmptyState'
import Loader from '../common/Loader'

export interface TableColumn<T> {
  key: keyof T | string
  header: string
  render?: (row: T) => ReactNode
}

export interface TableProps<T extends object> {
  columns: TableColumn<T>[]
  data: T[]
  emptyMessage?: string
  loading?: boolean
  className?: string
}

function getCellValue<T extends object>(row: T, key: keyof T | string): ReactNode {
  if (typeof key === 'string' && !(key in row)) {
    return null
  }
  const value = row[key as keyof T]
  if (value === null || value === undefined) {
    return null
  }
  if (typeof value === 'string' || typeof value === 'number' || typeof value === 'boolean') {
    return String(value)
  }
  return null
}

export default function Table<T extends object>({
  columns,
  data,
  emptyMessage = 'No data available.',
  loading = false,
  className = '',
}: TableProps<T>) {
  if (loading) {
    return <Loader label="Loading table data" />
  }

  if (data.length === 0) {
    return <EmptyState title={emptyMessage} />
  }

  return (
    <div className={`pl-table-wrap ${className}`.trim()}>
      <table className="pl-table">
        <thead>
          <tr>
            {columns.map((column) => (
              <th key={String(column.key)} scope="col">
                {column.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data.map((row, rowIndex) => (
            <tr key={rowIndex}>
              {columns.map((column) => (
                <td key={String(column.key)}>
                  {column.render
                    ? column.render(row)
                    : getCellValue(row, column.key)}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

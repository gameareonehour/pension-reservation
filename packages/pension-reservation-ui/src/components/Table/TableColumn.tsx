import styles from '@/styles/Table.module.css'
import { FC, ReactNode } from 'react'

type Props = {
  children: ReactNode
  width?: string
  colSpan?: number
}

const TableColumn: FC<Props> = ({ children, colSpan, width = 'auto' }) => {
  return (
    <th className={styles.table_column} style={{ width }} colSpan={colSpan}>
      {children}
    </th>
  )
}

export default TableColumn

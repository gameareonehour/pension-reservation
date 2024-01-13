import styles from '@/styles/Table.module.css'
import { FC, ReactNode } from 'react'

const TableCell: FC<{ children: ReactNode }> = ({ children }) => {
  return <td className={styles.table_cell}>{children}</td>
}

export default TableCell

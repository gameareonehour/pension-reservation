import { FC, ReactNode } from 'react'
import styles from './style.module.css'

const TableCell: FC<{ children: ReactNode }> = ({ children }) => {
  return <td className={styles.table_cell}>{children}</td>
}

export default TableCell

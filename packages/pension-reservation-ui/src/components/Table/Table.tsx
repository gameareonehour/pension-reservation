import styles from '@/styles/Table.module.css'
import { FC, ReactNode } from 'react'

const Table: FC<{ children: ReactNode }> = ({ children }) => {
  return <table className={styles.table}>{children}</table>
}

export default Table

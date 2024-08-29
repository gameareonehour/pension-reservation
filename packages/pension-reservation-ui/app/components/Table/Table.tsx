import { FC, ReactNode } from 'react'
import styles from './style.module.css'

const Table: FC<{ children: ReactNode }> = ({ children }) => {
  return <table className={styles.table}>{children}</table>
}

export default Table

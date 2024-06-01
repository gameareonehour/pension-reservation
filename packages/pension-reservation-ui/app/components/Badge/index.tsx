import { FC, ReactNode } from 'react'
import styles from './style.module.css'

const Badge: FC<{ children: ReactNode }> = ({ children }) => {
  return <div className={styles.badge}>{children}</div>
}

export default Badge

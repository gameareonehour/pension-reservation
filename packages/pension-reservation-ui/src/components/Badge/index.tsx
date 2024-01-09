import styles from '@/styles/Badge.module.css'
import { FC } from 'react'

const Badge: FC<{ children: ReactElement }> = ({ children }) => {
  return <div className={styles.badge}>{children}</div>
}

export default Badge

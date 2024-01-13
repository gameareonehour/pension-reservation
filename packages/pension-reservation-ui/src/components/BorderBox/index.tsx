import styles from '@/styles/BorderBox.module.css'
import { FC, ReactNode } from 'react'

const BorderBox: FC<{ children: ReactNode }> = ({ children }) => {
  return <div className={styles.border_box}>{children}</div>
}

export default BorderBox

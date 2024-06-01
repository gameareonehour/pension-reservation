import { FC, ReactNode } from 'react'
import styles from './style.module.css'

const BorderBox: FC<{ children: ReactNode }> = ({ children }) => {
  return <div className={styles.border_box}>{children}</div>
}

export default BorderBox

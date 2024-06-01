import { FC, ReactNode } from 'react'
import styles from './style.module.css'

const Footer: FC<{ children: ReactNode }> = ({ children }) => {
  return <div className={styles.footer}>{children}</div>
}

export default Footer

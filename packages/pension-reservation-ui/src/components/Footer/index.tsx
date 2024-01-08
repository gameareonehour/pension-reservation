import styles from '@/styles/Footer.module.css'
import { FC, ReactNode } from 'react'

const Footer: FC<{ children: ReactNode }> = ({ children }) => {
  return <div className={styles.footer}>{children}</div>
}

export default Footer

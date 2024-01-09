import styles from '@/styles/Footer.module.css'
import { FC } from 'react'

const Footer: FC<{ children: ReactElement }> = ({ children }) => {
  return <div className={styles.footer}>{children}</div>
}

export default Footer

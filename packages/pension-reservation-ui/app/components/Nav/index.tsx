import { FC, ReactNode } from 'react'
import styles from './style.module.css'

const Nav: FC<{ children: ReactNode }> = ({ children }) => {
  return <div className={styles.nav}>{children}</div>
}

export default Nav

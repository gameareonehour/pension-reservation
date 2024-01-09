import styles from '@/styles/Nav.module.css'
import { FC } from 'react'

const Nav: FC<{ children: ReactElement }> = ({ children }) => {
  return <div className={styles.nav}>{children}</div>
}

export default Nav

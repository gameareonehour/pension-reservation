import styles from '@/styles/Nav.module.css'
import { FC, ReactNode } from 'react'

const Nav: FC<{ children: ReactNode }> = ({ children }) => {
  return (
    <div className={styles.nav}>
      {children}
    </div>
  )    
}

export default Nav

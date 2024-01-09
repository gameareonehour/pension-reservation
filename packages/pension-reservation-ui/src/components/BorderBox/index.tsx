import styles from '@/styles/BorderBox.module.css'
import { FC } from 'react'

const BorderBox: FC<{ children: ReactElement }> = ({ children }) => {
  return <div className={styles.border_box}>{children}</div>
}

export default BorderBox

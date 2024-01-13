import styles from '@/styles/Button.module.css'
import { FC, ReactNode } from 'react'

type Props = {
  children: ReactNode
  onClick?: () => void
}

const Button: FC<Props> = ({ children, onClick }) => {
  return (
    <button className={styles.button} onClick={() => onClick?.()}>
      {children}
    </button>
  )
}

export default Button

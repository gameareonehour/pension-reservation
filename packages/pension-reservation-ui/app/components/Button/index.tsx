import { FC, ReactNode } from 'react'
import styles from './style.module.css'

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

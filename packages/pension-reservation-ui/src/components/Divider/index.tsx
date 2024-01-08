import styles from '@/styles/Divider.module.css'
import { FC } from 'react'

type Props = {
  margin?: string
}

const Divider: FC<Props> = ({ margin = '0px' }) => {
  return (
    <div className={styles.divider} style={{ margin: `${margin} 0` }}></div>
  )
}

export default Divider

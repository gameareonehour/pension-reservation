import { FC } from 'react'
import styles from './style.module.css'

type Props = {
  margin?: string
}

const Divider: FC<Props> = ({ margin = '0px' }) => {
  return <div className={styles.divider} style={{ margin: `${margin} 0` }}></div>
}

export default Divider

import styles from '@/styles/Logo.module.css'
import type { FC } from 'react'

const Logo: FC = () => {
  return (
    <div className={styles.logo}>
      <div className={styles.logo__annotation}>
        部活・サークル等のグループ利用に最適！アットホームなペンション！
      </div>

      <div className={styles.logo__mainText}>
        Marvelous Pension
      </div>

      <div className={styles.logo__subText}>
        天然温泉に入れるカジュアルペンション
      </div>
    </div>
  )
}

export default Logo

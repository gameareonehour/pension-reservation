import { redirect } from '@remix-run/node'
import type { FC } from 'react'
import { routes } from '~/core/routes'
import styles from './style.module.css'

const Logo: FC = () => {
  const navigateToReservationTop = () => {
    redirect(routes.top)
  }

  return (
    <div className={styles.logo} onClick={() => navigateToReservationTop()}>
      <div className={styles.logo__annotation}>
        部活・サークル等のグループ利用に最適！アットホームなペンション！
      </div>

      <div className={styles.logo__mainText}>Marvelous Pension</div>

      <div className={styles.logo__subText}>天然温泉に入れるカジュアルペンション</div>
    </div>
  )
}

export default Logo

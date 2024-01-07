import styles from '@/styles/components/Logo.module.css'

export default function Logo() {
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

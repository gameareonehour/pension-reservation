import styles from '@/styles/Menu.module.css'
import { FC } from 'react'

type Props = {
  items: Array<{
    name: string
    subMenus: Array<{
      name: string
      onClick: () => void
    }>
  }>
}

const Menu: FC<Props> = ({ items }) => {
  const inner = items.map((item, index) => {
    const subMenus = item.subMenus.map((subMenu, subMenuIndex) => {
      return (
        <div
          className={styles.menu__submenu_name}
          key={subMenuIndex}
          onClick={() => subMenu.onClick()}
        >
          <div className={styles.menu__submenu_name_text}>{subMenu.name}</div>
        </div>
      )
    })

    return (
      <span key={index}>
        <div className={styles.menu__name}>
          <div className={styles.menu__name_text}>{item.name}</div>
        </div>
        {subMenus}
      </span>
    )
  })

  return <div className={styles.menu}>{inner}</div>
}

export default Menu

import styles from '@/styles/Tabs.module.css'
import { FC } from 'react'

type Props = {
  items: Array<{
    name: string
    onClick: () => void
  }>
  selectedIndex?: number
}

const Tabs: FC<Props> = ({ items, selectedIndex }) => {
  const inner = items.map((item, index) => {
    return (
      <span key={index} onClick={() => item.onClick()}>
        <Tab isSelected={index === selectedIndex}>{item.name}</Tab>
      </span>
    )
  })

  return <div className={styles.tabs}>{inner}</div>
}

const Tab: FC<{ children: ReactElement; isSelected?: boolean }> = ({
  children,
  isSelected: selected,
}) => {
  const getTabTextClasses = (): string => {
    const classes = [styles.tab__text]

    if (selected) {
      classes.push(styles.tab__text_with_selected)
    }

    return classes.join(' ')
  }

  return (
    <div className={styles.tab}>
      <div className={getTabTextClasses()}>{children}</div>
    </div>
  )
}

export default Tabs

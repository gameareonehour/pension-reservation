import styles from '@/styles/Slot.module.css'
import type { FC, ReactNode } from 'react'

type Props = {
  children: ReactNode
  direction?: 'horizontal' | 'vertical'
  gap?:
    | 'none'
    | 'xs'
    | 'small'
    | 'middle'
    | 'medium'
    | 'large'
    | 'xl'
    | '2xl'
    | '3xl'
  align?: 'start' | 'center' | 'end' | ''
  justify?: 'start' | 'center' | 'end' | 'between' | 'around'
  wrap?: boolean
  fullWidth?: boolean
  fullHeight?: boolean
  width?: string
}

const Slot: FC<Props> = ({
  children,
  direction = 'horizontal',
  gap = 'none',
  align = 'start',
  justify = 'start',
  wrap = false,
  fullWidth = false,
  fullHeight = false,
  width = undefined,
}) => {
  const getClasses = (): string => {
    const classes = [styles.slot]

    switch (direction) {
      case 'horizontal':
        classes.push(styles.slot_with_horizontal)
        break
      case 'vertical':
        classes.push(styles.slot_with_vertical)
        break
    }

    switch (gap) {
      case 'none':
        classes.push(styles.slot_with_gap_none)
        break
      case 'xs':
        classes.push(styles.slot_with_gap_xs)
        break
      case 'small':
        classes.push(styles.slot_with_gap_small)
        break
      case 'middle':
        classes.push(styles.slot_with_gap_middle)
        break
      case 'medium':
        classes.push(styles.slot_with_gap_medium)
        break
      case 'large':
        classes.push(styles.slot_with_gap_large)
        break
      case 'xl':
        classes.push(styles.slot_with_gap_xl)
        break
      case '2xl':
        classes.push(styles.slot_with_gap_2xl)
        break
      case '3xl':
        classes.push(styles.slot_with_gap_3xl)
        break
    }

    switch (align) {
      case 'start':
        classes.push(styles.slot_with_align_start)
        break
      case 'center':
        classes.push(styles.slot_with_align_center)
        break
      case 'end':
        classes.push(styles.slot_with_align_end)
        break
    }

    switch (justify) {
      case 'start':
        classes.push(styles.slot_with_justify_start)
        break
      case 'center':
        classes.push(styles.slot_with_justify_center)
        break
      case 'end':
        classes.push(styles.slot_with_justify_end)
        break
      case 'between':
        classes.push(styles.slot_with_justify_between)
        break
      case 'around':
        classes.push(styles.slot_with_justify_around)
        break
    }

    if (wrap) {
      classes.push(styles.slot_with_wrap)
    }

    if (fullWidth) {
      classes.push(styles.slot_with_full_width)
    }

    if (fullHeight) {
      classes.push(styles.slot_with_full_height)
    }

    return classes.join(' ')
  }

  return (
    <div className={getClasses()} style={{ width }}>
      {children}
    </div>
  )
}

export default Slot

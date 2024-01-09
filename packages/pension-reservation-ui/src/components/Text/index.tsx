import styles from '@/styles/Text.module.css'
import { FC } from 'react'

type Props = {
  children: ReactElement
  color?: 'natural' | 'black' | 'primary'
  size?: 'base' | '3xl' | '4xl'
  family?: 'noto_sans_jp' | 'noto_sans_hebrew'
  bold?: boolean
  light?: boolean
  lineHeight?: string
}

const Text: FC<Props> = ({
  children,
  color = 'natural',
  size = 'base',
  family = 'noto_sans_jp',
  bold = false,
  light = false,
  lineHeight = '160%',
}) => {
  const getClasses = (): string => {
    const classes = [styles.text]

    switch (color) {
      case 'natural':
        classes.push(styles.text_with_natural)
        break
      case 'black':
        classes.push(styles.text_with_black)
        break
      case 'primary':
        classes.push(styles.text_with_primary)
        break
    }

    switch (size) {
      case 'base':
        classes.push(styles.text_with_size_base)
        break
      case '3xl':
        classes.push(styles.text_with_size_3xl)
        break
      case '4xl':
        classes.push(styles.text_with_size_4xl)
        break
    }

    switch (family) {
      case 'noto_sans_jp':
        classes.push(styles.text_with_noto_sans_jp)
        break
      case 'noto_sans_hebrew':
        classes.push(styles.text_with_noto_sans_jp)
        break
    }

    if (bold) {
      classes.push(styles.text_with_bold)
    }

    if (light) {
      classes.push(styles.text_with_light)
    }

    return classes.join(' ')
  }

  return (
    <div className={getClasses()} style={{ lineHeight }}>
      {children}
    </div>
  )
}

export default Text

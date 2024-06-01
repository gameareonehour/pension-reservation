import Image from 'next/image'
import { FC } from 'react'

export const LoadingSkeleton: FC = () => {
  return (
    <Image
      src={'/loading-skeleton-top-release-note.svg'}
      alt='hero image'
      width={1440}
      height={350}
      style={{ width: '100%', height: 'auto' }}
      priority
    />
  )
}

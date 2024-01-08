'use client'

import heroImage from '@/assets/images/hero-image.jpeg'
import Badge from '@/components/Badge'
import Footer from '@/components/Footer'
import Logo from '@/components/Logo'
import Menu from '@/components/Menu'
import Nav from '@/components/Nav'
import Slot from '@/components/Slot'
import Tabs from '@/components/Tabs'
import Text from '@/components/Text'
import { routes } from '@/core/routes'
import styles from '@/styles/Layout.module.css'
import Image from 'next/image'
import { usePathname, useRouter } from 'next/navigation'

import { useMemo, type FC, type ReactNode } from 'react'

const Layout: FC<{ children: ReactNode }> = ({ children }) => {
  const router = useRouter()
  const pathname = usePathname()

  const selectedTabIndex = useMemo((): number | undefined => {
    switch (pathname) {
      case routes.reservationTop:
        return 0
      case routes.roomCatalog:
        return 1
      case routes.vacancyRoomSearch:
        return 2
      default:
        return undefined
    }
  }, [pathname])

  const tabItems = [
    {
      name: 'トップ',
      onClick: () => {
        router.push(routes.reservationTop)
      },
    },
    {
      name: 'お部屋紹介',
      onClick: () => {
        router.push(routes.roomCatalog)
      },
    },
    {
      name: 'ご予約',
      onClick: () => {
        router.push(routes.vacancyRoomSearch)
      },
    },
  ]

  const menus = [
    {
      name: 'ご予約',
      subMenus: [
        {
          name: '空室検索',
          onClick: () => {
            router.push(routes.vacancyRoomSearch)
          },
        },
      ],
    },
    {
      name: 'お部屋紹介',
      subMenus: [
        {
          name: '和室',
          onClick: () => {
            router.push(routes.roomCatalog + '?type=1')
          },
        },
        {
          name: '洋室',
          onClick: () => {
            router.push(routes.roomCatalog + '?type=2')
          },
        },
        {
          name: '和洋室',
          onClick: () => {
            router.push(routes.roomCatalog + '?type=3')
          },
        },
      ],
    },
  ]

  return (
    <div className={styles.layout}>
      <Slot fullWidth direction={'vertical'} gap={'large'}>
        <Nav>
          <Logo />

          <Slot direction={'vertical'} gap={'small'}>
            <Badge>ご予約／お問い合わせ</Badge>
            <Text color={'black'} size={'3xl'} bold>
              ☎︎0120-000-000
            </Text>
          </Slot>
        </Nav>

        <Slot fullWidth direction={'vertical'}>
          <Tabs items={tabItems} selectedIndex={selectedTabIndex} />
          <Image
            src={heroImage}
            alt="hero image"
            style={{ width: '100%', height: 'auto' }}
            priority
          />
        </Slot>

        <Slot fullWidth gap={'large'}>
          <Menu items={menus} />

          {children}
        </Slot>
      </Slot>

      <div className={styles.layout__bottom_padding}></div>
      <Footer>
        Copyright © {new Date().getFullYear()} Marvelous Pension All Rights
        Reserved.
      </Footer>
    </div>
  )
}

export default Layout

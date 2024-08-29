// import { redirect } from '@remix-run/node'

import { useLocation, useNavigate } from '@remix-run/react'
import { useMemo, type FC, type ReactNode } from 'react'
import Badge from '~/components/Badge'
import Footer from '~/components/Footer'
import Logo from '~/components/Logo'
import Menu from '~/components/Menu'
import Nav from '~/components/Nav'
import Slot from '~/components/Slot'
import Tabs from '~/components/Tabs'
import Text from '~/components/Text'
import { routes } from '~/core/routes'
import styles from './style.module.css'

type Props = {
  children: ReactNode
  heroImageURI?: string
}

const Layout: FC<Props> = ({ children, heroImageURI }) => {
  const location = useLocation()
  const navigate = useNavigate()

  const selectedTabIndex = useMemo((): number | undefined => {
    switch (location.pathname) {
      case routes.top:
        return 0
      case routes.catalog:
        return 1
      case routes.vacancyRoomSearch:
        return 2
      default:
        // カタログ詳細画面の場合
        const regexp = new RegExp(`${routes.catalogDetails}\\d`)
        if (regexp.test(location.pathname)) {
          return 1
        }

        return undefined
    }
  }, [location.pathname])

  const tabItems = [
    {
      name: 'トップ',
      onClick: () => {
        navigate(routes.top)
      },
    },
    {
      name: 'お部屋紹介',
      onClick: () => {
        navigate(routes.catalog)
      },
    },
    {
      name: 'ご予約',
      onClick: () => {
        navigate(routes.vacancyRoomSearch)
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
            navigate(routes.vacancyRoomSearch)
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
            console.log('呼ばれたど')
            navigate(routes.catalog + '?type=1')
          },
        },
        {
          name: '洋室',
          onClick: () => {
            navigate(routes.catalog + '?type=2')
          },
        },
        {
          name: '和洋室',
          onClick: () => {
            navigate(routes.catalog + '?type=3')
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
          {heroImageURI && (
            <img
              src={heroImageURI}
              alt='hero image'
              width={787}
              height={350}
              style={{ width: '100%', height: 'auto' }}
            />
          )}
        </Slot>

        <Slot fullWidth gap={'large'}>
          <Menu items={menus} />

          {children}
        </Slot>
      </Slot>

      <div className={styles.layout__bottom_padding}></div>
      <Footer>
        Copyright © {new Date().getFullYear()} Marvelous Pension All Rights Reserved.
      </Footer>
    </div>
  )
}

export default Layout

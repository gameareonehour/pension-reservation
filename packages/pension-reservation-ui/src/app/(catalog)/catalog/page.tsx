'use client'

import { GetCatalogData } from '@/app/(catalog)/_lib/types'
import BorderBox from '@/components/BorderBox'
import Button from '@/components/Button'
import Divider from '@/components/Divider'
import Layout from '@/components/Layout'
import Slot from '@/components/Slot'
import Table from '@/components/Table/Table'
import TableBody from '@/components/Table/TableBody'
import TableCell from '@/components/Table/TableCell'
import TableColumn from '@/components/Table/TableColumn'
import TableHead from '@/components/Table/TableHead'
import TableRow from '@/components/Table/TableRow'
import Text from '@/components/Text'
import { routes } from '@/core/routes'
import { fetcher } from '@/core/swr/fetcher'
import Image from 'next/image'
import { useRouter, useSearchParams } from 'next/navigation'
import { useMemo } from 'react'
import useSWR from 'swr'

export default function RoomCatalog() {
  const router = useRouter()
  const query = useSearchParams()
  const type = query.get('type') ?? ''
  const { data, isLoading } = useSWR<GetCatalogData>(`/api/catalog?type=${type}`, fetcher)

  const rooms = useMemo(() => {
    const navigateToCatalogDetails = (roomId: number): void => {
      router.push(routes.roomCatalogDetails + roomId)
    }

    if (isLoading) {
      return (
        <Image
          src={'/loading-skeleton-catalog-rooms.svg'}
          alt='hero image'
          width={1440}
          height={350}
          style={{ width: '100%', height: 'auto' }}
          priority
        />
      )
    }

    return (
      <Table>
        <TableHead>
          <TableRow>
            <TableColumn width='230px'>お部屋名称</TableColumn>
            <TableColumn>お部屋タイプ</TableColumn>
            <TableColumn>
              一泊料金
              <br />
              （部屋単位）
            </TableColumn>
            <TableColumn colSpan={2}>お部屋イメージ</TableColumn>
          </TableRow>
        </TableHead>
        <TableBody>
          {data?.items.map((item, index) => {
            return (
              <TableRow key={index}>
                <TableCell>{item.name}</TableCell>
                <TableCell>{item.type}</TableCell>
                <TableCell>{item.dayfee}</TableCell>
                <TableCell>
                  <Image src={item.imageUrl} alt='room image' width={320} height={240} priority />
                </TableCell>
                <TableCell>
                  <Button onClick={() => navigateToCatalogDetails(item.id)}>詳細</Button>
                </TableCell>
              </TableRow>
            )
          })}
        </TableBody>
      </Table>
    )
  }, [data?.items, isLoading, router])

  return (
    <Layout>
      <Slot direction={'vertical'} fullWidth gap={'large'}>
        <Slot direction={'vertical'} fullWidth gap={'medium'}>
          <BorderBox>
            <Text color={'natural'}>お部屋の紹介</Text>
          </BorderBox>
          <Slot fullWidth direction={'vertical'}>
            <Text color={'natural'} bold>
              自慢のお部屋をご紹介
            </Text>

            <Divider margin='8px' />

            <Text color={'natural'}>
              和室・洋室・和洋室と、ご希望に沿った形でお部屋をお選び頂けます。
            </Text>
          </Slot>
          {rooms}
        </Slot>
      </Slot>
    </Layout>
  )
}

import { paths } from '@/core/api/openapi'

export type GetCatalogResponse =
  paths['/catalog']['get']['responses']['200']['content']['application/json']

export type GetCatalogData = {
  items: Array<{
    id: number
    name: string
    type: string
    dayfee: string
    imageUrl: string
  }>
}

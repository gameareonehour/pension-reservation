import { paths } from '@/core/api/openapi'

export type GetLatestReleaseNotesResponse =
  paths['/release-notes']['get']['responses']['200']['content']['application/json']

export type GetLatestReleaseNotesData = {
  items: Array<{
    id: number
    text: string
    createdAt: string
  }>
}

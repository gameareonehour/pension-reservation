import { GetLatestReleaseNotesData, GetLatestReleaseNotesResponse } from './types'

export const responseToReleaseNotes = (
  response: GetLatestReleaseNotesResponse
): GetLatestReleaseNotesData => {
  const formatter = new Intl.DateTimeFormat('fr-CA', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
  })

  const items = response.items.map((item) => {
    const createdAt = formatter.format(new Date(item.created_at))

    return {
      id: item.id,
      text: item.text,
      createdAt,
    }
  })

  return { items }
}

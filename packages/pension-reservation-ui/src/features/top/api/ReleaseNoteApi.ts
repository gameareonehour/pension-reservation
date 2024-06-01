import { GetLatestReleaseNotesData, GetLatestReleaseNotesResponse } from './types'

export class ReleaseNoteApi {
  private domain: string

  constructor(domain: string) {
    this.domain = domain
  }

  async getLatestReleaseNotes(): Promise<GetLatestReleaseNotesData> {
    const response = await fetch(`${this.domain}/api/v1/release-notes`, {
      next: { revalidate: 60 },
    })

    const body = (await response.json()) as GetLatestReleaseNotesResponse

    const formatter = new Intl.DateTimeFormat('fr-CA', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    })

    const items = body.items.map((item) => {
      const createdAt = formatter.format(new Date(item.created_at))

      return {
        id: item.id,
        text: item.text,
        createdAt,
      }
    })

    return { items }
  }
}

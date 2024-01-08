import { GetLatestReleaseNotesResponse } from './types'

export class ReleaseNoteApi {
  private domain: string

  constructor(domain: string) {
    this.domain = domain
  }

  async getLatestReleaseNotes(): Promise<GetLatestReleaseNotesResponse> {
    const response = await fetch(`${this.domain}/api/v1/release-notes`, {
      next: { revalidate: 60 },
    })

    return response.json()
  }
}

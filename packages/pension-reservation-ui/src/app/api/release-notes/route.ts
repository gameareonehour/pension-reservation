import { domain } from '@/core/api/configuration'
import { ReleaseNoteApi } from '@/features/releaseNotes/api'
import { responseToReleaseNotes } from '@/features/releaseNotes/transfer'
import { NextResponse } from 'next/server'

export async function GET() {
  const api = new ReleaseNoteApi(domain)

  const response = await api.getLatestReleaseNotes()
  const data = responseToReleaseNotes(response)

  return NextResponse.json(data)
}

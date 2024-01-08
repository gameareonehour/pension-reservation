import { ReleaseNoteApi } from '@/app/(top)/_lib/api'
import { responseToReleaseNotes } from '@/app/(top)/_lib/transfer'
import { domain } from '@/core/api/configuration'
import { NextResponse } from 'next/server'

export async function GET() {
  const api = new ReleaseNoteApi(domain)

  const response = await api.getLatestReleaseNotes()
  const data = responseToReleaseNotes(response)

  return NextResponse.json(data)
}

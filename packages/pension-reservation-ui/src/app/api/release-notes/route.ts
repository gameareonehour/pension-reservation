import { ReleaseNoteApi } from '@/app/(top)/_lib/api'
import { domain } from '@/core/api/configuration'
import { NextResponse } from 'next/server'

export async function GET() {
  const api = new ReleaseNoteApi(domain)
  const data = await api.getLatestReleaseNotes()

  return NextResponse.json(data)
}

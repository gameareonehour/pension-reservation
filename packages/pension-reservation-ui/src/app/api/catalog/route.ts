import { CatalogApi } from '@/app/(catalog)/_lib/api'
import { domain } from '@/core/api/configuration'
import { NextRequest, NextResponse } from 'next/server'

export async function GET(req: NextRequest) {
  const query = req.nextUrl.searchParams
  const type = query.get('type')

  const api = new CatalogApi(domain)
  const data = await api.getCatalog(type)

  return NextResponse.json(data)
}

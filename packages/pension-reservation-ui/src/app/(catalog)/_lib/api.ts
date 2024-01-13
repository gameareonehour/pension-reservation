import { formatMoney } from '@/core/shared/format'
import { GetCatalogData, GetCatalogResponse } from './types'

export class CatalogApi {
  private domain: string

  constructor(domain: string) {
    this.domain = domain
  }

  async getCatalog(type: string | null): Promise<GetCatalogData> {
    let url = ''
    if (type) {
      url = `${this.domain}/api/v1/catalog?type=${type}`
    } else {
      url = `${this.domain}/api/v1/catalog`
    }

    const response = await fetch(url, {
      next: { revalidate: 60 },
    })

    const body = (await response.json()) as GetCatalogResponse

    const items = body.items.map((item) => {
      return {
        id: item.id,
        name: item.name,
        type: item.type,
        dayfee: formatMoney(item.dayfee),
        imageUrl: item.image_url,
      }
    })

    return { items }
  }
}

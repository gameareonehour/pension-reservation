import '@/styles/_globals.css'
import '@/styles/_variables.css'
import type { Metadata } from 'next'
import { Noto_Sans_Hebrew, Noto_Sans_JP, Orelega_One } from 'next/font/google'

const notoSansJp = Noto_Sans_JP<'--font-NotoSansJP'>({
  weight: ['100', '200', '300', '400', '500', '600', '700'],
  style: ['normal'],
  subsets: ['latin'],
  variable: '--font-NotoSansJP'
})

const notoSansHebrew = Noto_Sans_Hebrew<'--font-NotoSansHebrew'>({
  weight: ['300', '400'],
  style: ['normal'],
  subsets: ['latin'],
  variable: '--font-NotoSansHebrew'
})

const orelegaOne = Orelega_One<'--font-OrelegaOne'>({
  weight: '400',
  style: ['normal'],
  subsets: ['latin'],
  variable: '--font-OrelegaOne'
})

export const metadata: Metadata = { title: 'ペンション予約サイト' }

export default function RootLayout({children}: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body className={`${notoSansJp.className} ${notoSansJp.variable} ${notoSansHebrew.variable} ${orelegaOne.variable}`} style={{ width: "100%" }}>
        {children}
      </body>
    </html>
  )
}

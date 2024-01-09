import BorderBox from '@/components/BorderBox'
import Divider from '@/components/Divider'
import Layout from '@/components/Layout'
import Slot from '@/components/Slot'
import Text from '@/components/Text'

export default function RoomCatalog() {
  return (
    <Layout>
      <Slot direction={'vertical'} fullWidth gap={'large'}>
        <Slot direction={'vertical'} fullWidth gap={'medium'}>
          <BorderBox>
            <Text color={'natural'}>お部屋の紹介</Text>
          </BorderBox>
          <Slot fullWidth direction={'vertical'}>
            <Text color={'natural'} bold>
              にこやか笑顔でお出迎え
            </Text>

            <Divider margin='8px' />

            <Text color={'natural'}>
              和室・洋室・和洋室と、ご希望に沿った形でお部屋をお選び頂けます。
            </Text>
          </Slot>
          a
        </Slot>
      </Slot>
    </Layout>
  )
}

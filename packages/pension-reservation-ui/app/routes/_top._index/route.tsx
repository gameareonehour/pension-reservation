import BorderBox from '~/components/BorderBox'
import Divider from '~/components/Divider'
import Layout from '~/components/Layout'
import Slot from '~/components/Slot'
import Text from '~/components/Text'

export default function Index() {
  const releaseNotes: any[] = []

  return (
    <Layout heroImageURI='/hero-image.jpeg'>
      <Slot fullWidth direction={'vertical'} gap={'large'}>
        <Slot fullWidth direction={'vertical'} gap={'middle'}>
          <BorderBox>
            <Slot direction={'vertical'} gap={'xs'}>
              <Text
                size={'4xl'}
                family={'noto_sans_hebrew'}
                color={'primary'}
                lineHeight='160%'
                light
              >
                What’s New
              </Text>

              <Text color={'natural'}>更新情報</Text>
            </Slot>
          </BorderBox>

          <Slot direction={'vertical'} gap={'small'}>
            {releaseNotes}
          </Slot>
        </Slot>

        <Slot fullWidth direction={'vertical'} gap={'middle'}>
          <BorderBox>
            <Slot direction={'vertical'} gap={'xs'}>
              <Text
                size={'4xl'}
                family={'noto_sans_hebrew'}
                color={'primary'}
                lineHeight='160%'
                light
              >
                About Us
              </Text>

              <Text color={'natural'}>Marvelous Pensionについて</Text>
            </Slot>
          </BorderBox>

          <Slot fullWidth direction={'vertical'} gap={'middle'}>
            <Slot fullWidth direction={'vertical'}>
              <Text color={'natural'} bold>
                にこやか笑顔でお出迎え
              </Text>

              <Divider margin='8px' />

              <Text color={'natural'}>
                少人数で営業している当ペンションですが、スタッフ一同心掛けているのは、
                <br />
                にこやかな笑顔で接客対応することです！
              </Text>
            </Slot>

            <Slot fullWidth direction={'vertical'}>
              <Text color={'natural'} bold>
                大人数でワイワイと
              </Text>

              <Divider margin='8px' />

              <Text color={'natural'}>
                部活やサークル、仲の良いお友達同士などと、
                <br />
                大人数でワイワイしながら過ごすのに最適な環境づくりを目指しています！
              </Text>
            </Slot>

            <Slot fullWidth direction={'vertical'}>
              <Text color={'natural'} bold>
                観光スポットに恵まれた好立地
              </Text>

              <Divider margin='8px' />

              <Text color={'natural'}>
                ゲレンデ、テニスコート、各種レクリエーション施設へのアクセスは抜群です！
                <br />
                また、温泉地の中心街からも近いため、観光にも最適です！
              </Text>
            </Slot>
          </Slot>
        </Slot>
      </Slot>
    </Layout>
  )
}

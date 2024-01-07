import Badge from '@/components/Badge';
import Logo from '@/components/Logo';
import Nav from '@/components/Nav';
import Slot from '@/components/Slot';
import Text from '@/components/Text';
import styles from '@/styles/Layout.module.css';
import type { FC, ReactNode } from "react";

const Layout: FC<{ children: ReactNode }> = ({ children }) => {
	return (
	  <div className={styles.layout}>
      <Slot fullWidth gap={'large'}>
        <Nav>
          <Logo />

          <Slot direction={'vertical'} gap={'small'}>
            <Badge>
              ご予約／お問い合わせ
            </Badge>
            <Text color={'black'} size={'3xl'} bold>
              ☎︎0120-000-000
            </Text>
          </Slot>
        </Nav>

        
      </Slot>

      {children}
    </div>
	)
}

export default Layout

"use client"

import heroImage from "@/assets/images/hero-image.jpeg";
import Badge from "@/components/Badge";
import Logo from "@/components/Logo";
import Nav from "@/components/Nav";
import Slot from "@/components/Slot";
import Tabs from "@/components/Tabs";
import Text from "@/components/Text";
import { routes } from "@/core/routes";
import styles from "@/styles/Layout.module.css";
import Image from "next/image";
import { usePathname, useRouter } from "next/navigation";

import { useMemo, type FC, type ReactNode } from "react";

const Layout: FC<{ children: ReactNode }> = ({ children }) => {
  const router = useRouter()
  const pathname = usePathname()

  const selectedTabIndex = useMemo((): number | undefined => {
    switch (pathname) {
      case routes.reservationTop:
        return 0
      case routes.roomCatalog:
        return 1
      case routes.vacancyRoomSearch:
        return 2
      default:
        return undefined
    }
  }, [pathname]); 

  const tabItems = [
    { 
      name: "トップ",
      onClick: () => {
        router.push(routes.reservationTop);
      }
    },
    { 
      name: "お部屋紹介",
      onClick: () => {
        router.push(routes.roomCatalog);
      }
    },
    {
      name: "ご予約",
      onClick: () => {
        router.push(routes.vacancyRoomSearch);
      }
    },
  ];

  return (
    <div className={styles.layout}>
      <Slot fullWidth direction={"vertical"} gap={"large"}>
        <Nav>
          <Logo />

          <Slot direction={"vertical"} gap={"small"}>
            <Badge>ご予約／お問い合わせ</Badge>
            <Text color={"black"} size={"3xl"} bold>
              ☎︎0120-000-000
            </Text>
          </Slot>
        </Nav>

        <Slot fullWidth direction={"vertical"}>
          <Tabs items={tabItems} selectedIndex={selectedTabIndex} />
          <Image src={heroImage} alt="hero image" layout="responsive" />
        </Slot>
      </Slot>

      {children}
    </div>
  );
};

export default Layout;

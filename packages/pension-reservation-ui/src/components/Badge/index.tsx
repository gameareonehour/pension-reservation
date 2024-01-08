import styles from "@/styles/Badge.module.css";
import { FC, ReactNode } from "react";

const Badge: FC<{ children: ReactNode }> = ({ children }) => {
  return <div className={styles.badge}>{children}</div>;
};

export default Badge;

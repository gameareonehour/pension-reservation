import styles from "@/styles/Text.module.css";
import { FC, ReactNode } from "react";

type Props = {
  children: ReactNode;
  color?: "natural" | "black";
  size?: "base" | "3xl";
  bold?: boolean;
};

const Text: FC<Props> = ({
  children,
  color = "natural",
  size = "base",
  bold = false,
}) => {
  const getClasses = (): string => {
    const classes = [styles.text];

    switch (color) {
      case "natural":
        classes.push(styles.text_with_natural);
        break;
      case "black":
        classes.push(styles.text_with_black);
        break;
    }

    switch (size) {
      case "base":
        classes.push(styles.text_with_size_base);
        break;
      case "3xl":
        classes.push(styles.text_with_size_3xl);
        break;
    }

    if (bold) {
      classes.push(styles.text_with_bold);
    }

    return classes.join(" ");
  };

  return <div className={getClasses()}>{children}</div>;
};

export default Text;

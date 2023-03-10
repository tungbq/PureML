import { tv, type VariantProps } from "tailwind-variants";
import { Link } from "@remix-run/react";

const linkStyles = tv({
  base: "text-brand-200 hover:text-blue-250 font-medium text-brand-200",
  variants: {
    intent: {
      primary: "text-base w-fit h-fit",
      secondary: "",
    },
  },
  defaultVariants: {
    intent: "primary",
  },
});

interface Props extends VariantProps<typeof linkStyles> {
  [children: string]: any;
  hyperlink: string;
}

export default function link(props: Props) {
  return (
    <Link to={props.hyperlink} className={linkStyles(props)}>
      {props.children}
    </Link>
  );
}

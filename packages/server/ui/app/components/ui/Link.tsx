import { Link } from "@remix-run/react";
import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";

const linkStyles = cva("", {
  variants: {
    intent: {
      primary:
        "text-base text-blue-700 hover:text-blue-700 font-medium w-fit h-fit",
      secondary: "text-sm text-brand-200 hover:text-brand-300 font-medium",
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

export default function link({ intent, hyperlink, children }: Props) {
  return (
    <Link to={hyperlink} className={linkStyles({ intent })}>
      {children}
    </Link>
  );
}

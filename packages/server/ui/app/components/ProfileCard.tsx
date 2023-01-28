import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";

const cardStyles = cva(
  "items-center justify-center p-8 text-lg font-medium focus:outline-none",
  {
    variants: {
      intent: {
        profile: "bg-slate-0 text-slate-800 rounded-xl border border-slate-200",
      },
      fullWidth: {
        true: "w-full",
      },
    },
    defaultVariants: {
      intent: "profile",
      fullWidth: true,
    },
  }
);
type CardProps = VariantProps<typeof cardStyles>;

interface Props extends CardProps {
  count: string;
  title: string;
}

export default function ProfileCard({
  intent,
  fullWidth,
  count,
  title,
}: Props) {
  return (
    <div>
      <div className={cardStyles({ intent, fullWidth })}>
        <div className="text-3xl">{count}</div>
        <div className="text-lg">{title}</div>
      </div>
    </div>
  );
}

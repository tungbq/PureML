import { tv, type VariantProps } from "tailwind-variants";

const cardStyles = tv({
  base: "items-center justify-center p-8 text-lg font-medium focus:outline-none",
  variants: {
    intent: {
      profile: "bg-slate-0 text-slate-600 rounded-xl border border-slate-200",
    },
    fullWidth: {
      true: "w-full",
    },
  },
  defaultVariants: {
    intent: "profile",
    fullWidth: true,
  },
});
type CardProps = VariantProps<typeof cardStyles>;

interface Props extends CardProps {
  count: string;
  title: string;
}

export default function ProfileCard(props: Props) {
  return (
    <div>
      <div className={cardStyles(props)}>
        <div className="text-3xl">{props.count}</div>
        <div className="text-lg">{props.title}</div>
      </div>
    </div>
  );
}

import { tv, type VariantProps } from "tailwind-variants";

const buttonStyles = tv({
  base: "font-medium",
  variants: {
    intent: {
      primary: "btn btn-primary btn-sm text-white",
      secondary:
        "btn btn-secondary btn-sm bg-transparent text-slate-600 border border-slate-400",
      danger: "btn btn-error btn-sm",
      icon: "btn btn-secondary btn-sm bg-transparent text-slate-600 border border-slate-400",
      org: "btn btn-secondary text-slate-600",
    },
    fullWidth: {
      true: "w-full",
      false: "w-inherit",
    },
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: true,
  },
});

interface Props extends VariantProps<typeof buttonStyles> {
  [children: string]: any;
}

function Button(props: Props) {
  return (
    <button onClick={props.onClick} className={buttonStyles(props)}>
      {props.intent !== "icon" ? (
        <div className="flex items-center w-max">{props.children}</div>
      ) : (
        <div className="flex items-center w-max">{props.children}</div>
      )}
    </button>
  );
}

export default Button;

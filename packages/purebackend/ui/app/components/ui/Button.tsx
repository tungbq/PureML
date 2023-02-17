import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";

const buttonStyles = cva(
  "flex items-center font-medium focus:outline-none rounded-lg",
  {
    variants: {
      intent: {
        primary:
          "bg-blue-750 text-white px-4 py-2 h-9 hover:bg-blue-600 justify-center",
        landingpg:
          "bg-blue-750 text-white px-12 py-2 h-10 hover:bg-blue-600 justify-center",
        secondary:
          "bg-white text-slate-600 border border-slate-400 px-4 py-2 h-9 hover:bg-slate-100 justify-center",
        danger:
          "bg-red-400 text-white hover:bg-red-500 justify-center px-4 py-2 h-9",
        icon: "px-4 py-2 border border-slate-400 hover:bg-slate-100 text-slate-600",
        org: "px-4 py-2 hover:bg-slate-100 hover:text-slate-600",
      },
      fullWidth: {
        true: "w-full",
      },
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: true,
    },
  }
);

interface Props extends VariantProps<typeof buttonStyles> {
  [children: string]: any;
  icon: any;
}

function Button({ intent, fullWidth, children, icon, onClick }: Props) {
  return (
    <button onClick={onClick} className={buttonStyles({ intent, fullWidth })}>
      {intent !== "icon" ? (
        <div className="w-max">{children}</div>
      ) : (
        <div className="flex items-center">
          {icon}
          <div className="px-4 w-max">{children}</div>
        </div>
      )}
    </button>
  );
}

export default Button;

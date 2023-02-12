import * as Avatar from "@radix-ui/react-avatar";
import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";

const avatarStyles = cva(
  "flex items-center px-3 py-2 font-medium focus:outline-none",
  {
    variants: {
      intent: {
        primary:
          "w-6 h-6 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-full capitalize",
        profile: "text-black h-9 rounded justify-center capitalize",
        org: "bg-brand-100 text-black h-9 rounded-full justify-center capitalize",
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
interface Props extends VariantProps<typeof avatarStyles> {
  [children: string]: any;
}

export default function AvatarIcon({ intent, fullWidth, children }: Props) {
  return (
    <div>
      {intent === "primary" ? (
        <div className="h-full">
          <Avatar.Root>
            <Avatar.Fallback className={avatarStyles({ intent, fullWidth })}>
              {children}
            </Avatar.Fallback>
          </Avatar.Root>
        </div>
      ) : (
        <>
          {intent === "profile" ? (
            <div className="px-1">
              <Avatar.Root>
                <Avatar.Fallback
                  className={avatarStyles({ intent, fullWidth })}
                >
                  {/* className={buttonStyles({ intent, fullWidth })} */}
                  {children}
                </Avatar.Fallback>
              </Avatar.Root>
            </div>
          ) : (
            <div className="h-full">
              <Avatar.Root>
                <Avatar.Fallback
                  className={avatarStyles({ intent, fullWidth })}
                >
                  {children}
                </Avatar.Fallback>
              </Avatar.Root>
            </div>
          )}
        </>
      )}
    </div>
  );
}

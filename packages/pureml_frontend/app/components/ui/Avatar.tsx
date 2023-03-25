import * as Avatar from "@radix-ui/react-avatar";
import { tv, type VariantProps } from "tailwind-variants";

const avatarStyles = tv({
  base: "avatar flex items-center justify-center text-blue-600 rounded-full",
  variants: {
    intent: {
      primary: "bg-blue-200 w-6",
      profile: "w-6",
      org: "bg-blue-200 w-9 h-9",
      avatar: "bg-blue-200 w-8 h-8 text-base",
    },
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: true,
  },
});

interface Props extends VariantProps<typeof avatarStyles> {
  [children: string]: any;
}

export default function AvatarIcon(props: Props) {
  return (
    <div>
      {props.intent === "primary" ? (
        <div className="h-full">
          <Avatar.Root>
            <Avatar.Fallback className={avatarStyles(props)}>
              {props.children}
            </Avatar.Fallback>
          </Avatar.Root>
        </div>
      ) : (
        <>
          {props.intent === "profile" ? (
            <div className="px-1">
              <Avatar.Root>
                <Avatar.Fallback className={avatarStyles(props)}>
                  {props.children}
                </Avatar.Fallback>
              </Avatar.Root>
            </div>
          ) : (
            <div className="h-full">
              <Avatar.Root>
                <Avatar.Fallback className={avatarStyles(props)}>
                  {props.children}
                </Avatar.Fallback>
              </Avatar.Root>
            </div>
          )}
        </>
      )}
    </div>
  );
}

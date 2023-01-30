import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";
import * as DropdownMenu from "@radix-ui/react-dropdown-menu";
import { ChevronDown } from "lucide-react";
import { useNavigate } from "@remix-run/react";

const dropdownTrigger = cva("focus:outline-none z-50", {
  variants: {
    intent: {
      primary: "flex justify-between items-center z-50",
      orgType:
        "flex justify-between items-center text-sm text-slate-600 rounded border border-slate-600 h-8 px-2 z-50",
      branch:
        "flex justify-between items-center text-sm text-slate-600 rounded border border-slate-600 h-8 px-2 z-50",
    },
    fullWidth: {
      true: "w-full",
      false: "w-28",
    },
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: true,
  },
});
type TriggerProps = VariantProps<typeof dropdownTrigger>;

const dropdownContent = cva("focus:outline-none", {
  variants: {
    color: {
      primary:
        "bg-slate-100 justify-center items-center text-sm text-slate-600 rounded border border-slate-200 z-50 shadow",
    },
    contentWidth: {
      true: "w-full",
    },
  },
  defaultVariants: {
    color: "primary",
    contentWidth: true,
  },
});
type ContentProps = VariantProps<typeof dropdownContent>;

const dropdownItems = cva("focus:outline-none", {
  variants: {
    space: {
      primary:
        "flex px-3 py-2 text-sm text-base justify-left items-center rounded outline-none hover:bg-slate-200 cursor-pointer",
    },
  },
  defaultVariants: {
    space: "primary",
  },
});
type ItemProps = VariantProps<typeof dropdownItems>;

export interface DropdownProps extends TriggerProps, ContentProps, ItemProps {
  [children: string]: any;
}

export default function Dropdown({
  intent,
  fullWidth,
  color,
  contentWidth,
  space,
  children,
}: DropdownProps) {
  const navigate = useNavigate();

  return (
    <div>
      <DropdownMenu.Root>
        <DropdownMenu.Trigger
          className={dropdownTrigger({ intent, fullWidth })}
        >
          {children}
          {intent !== "primary" ? (
            <ChevronDown className="text-slate-400" />
          ) : (
            ""
          )}
        </DropdownMenu.Trigger>
        {intent === "primary" ? (
          <DropdownMenu.Content
            sideOffset={7}
            align="end"
            className={dropdownContent({ color, contentWidth })}
          >
            <DropdownMenu.Item
              className={dropdownItems({ space })}
              onClick={() => {
                navigate(`/profile`);
              }}
            >
              Profile
            </DropdownMenu.Item>
            <DropdownMenu.Item
              className={dropdownItems({ space })}
              onClick={() => {
                navigate(`/settings`);
              }}
            >
              Settings
            </DropdownMenu.Item>
            <DropdownMenu.Item
              className={dropdownItems({ space })}
              onClick={() => {
                navigate(`/contact`);
              }}
            >
              Contact Us
            </DropdownMenu.Item>
            <DropdownMenu.Item
              className={dropdownItems({ space })}
              onClick={() => {
                navigate(`/logout`);
              }}
            >
              Sign out
            </DropdownMenu.Item>
          </DropdownMenu.Content>
        ) : (
          <>
            {intent === "orgType" ? (
              <DropdownMenu.Content
                sideOffset={7}
                className={dropdownContent({ color, contentWidth })}
              >
                <DropdownMenu.Item className={dropdownItems({ space })}>
                  Company
                </DropdownMenu.Item>
                <DropdownMenu.Item className={dropdownItems({ space })}>
                  University
                </DropdownMenu.Item>
                <DropdownMenu.Item className={dropdownItems({ space })}>
                  Classroom
                </DropdownMenu.Item>
                <DropdownMenu.Item className={dropdownItems({ space })}>
                  Non-profit
                </DropdownMenu.Item>
                <DropdownMenu.Item className={dropdownItems({ space })}>
                  Community
                </DropdownMenu.Item>
              </DropdownMenu.Content>
            ) : (
              <DropdownMenu.Content
                sideOffset={7}
                className={dropdownContent({ color, contentWidth })}
              >
                {/* <DropdownMenu.Item className={dropdownItems({ space })}>
                  main
                </DropdownMenu.Item>
                <DropdownMenu.Item className={dropdownItems({ space })}>
                  dev
                </DropdownMenu.Item> */}
              </DropdownMenu.Content>
            )}
          </>
        )}
      </DropdownMenu.Root>
    </div>
  );
}

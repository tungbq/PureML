/* eslint-disable no-unused-vars */
import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";
import type { ChangeEvent } from "react";
import { Search } from "lucide-react";

const inputStyles = cva(
  "flex items-center justify-center px-4 py-2 focus:outline-none",
  {
    variants: {
      intent: {
        primary:
          "bg-slate-100 text-slate-600 rounded h-9 border-slate-600 border hover:border-blue-750 focus:border-blue-750",
        search:
          "bg-transparent text-sm text-slate-400 !justify-start rounded border-slate-200 border hover:border-blue-750 focus:border-blue-750",
        read: "bg-slate-100 text-slate-900 rounded h-9 border-slate-600 border hover:border-blue-750 focus:border-blue-750",
      },
      fullWidth: {
        true: "w-full",
        false: "w-[18rem]",
      },
      type: {
        text: "text",
        password: "password",
        email: "email",
        number: "number",
      },
    },
    defaultVariants: {
      intent: "primary",
      fullWidth: true,
      type: "text",
    },
  }
);

interface Props extends VariantProps<typeof inputStyles> {
  // [children: string]: any;
  placeholder?: string;
  ariaLabel?: string;
  dataTestid?: string;
  required?: boolean;
  name?: string;
  onChange?: (e: ChangeEvent<HTMLInputElement>) => void;
}

function Input({
  intent,
  fullWidth,
  type,
  name,
  placeholder,
  onChange,
  ariaLabel,
  dataTestid,
  required,
}: Props) {
  return (
    <div>
      {intent === "read" ? (
        <input
          required={required}
          type={type as string}
          className={inputStyles({ intent, fullWidth })}
          onChange={onChange}
          aria-label={ariaLabel}
          data-testid={dataTestid}
          defaultValue={placeholder}
          disabled
        />
      ) : (
        <div>
          {intent === "search" ? (
            <div className={inputStyles({ intent, fullWidth })}>
              <Search className="w-4 h-4" />
              <input
                required={required}
                type={type as string}
                className="border-none focus:outline-none pl-2 w-full"
                placeholder={placeholder}
                onChange={onChange}
                aria-label={ariaLabel}
                data-testid={dataTestid}
              />
            </div>
          ) : (
            <>
              <input
                required={required}
                type={type as string}
                name={name}
                className={inputStyles({ intent, fullWidth })}
                placeholder={placeholder}
                onChange={onChange}
                aria-label={ariaLabel}
                data-testid={dataTestid}
              />
            </>
          )}
        </div>
      )}
    </div>
  );
}

export default Input;

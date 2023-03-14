import { tv, type VariantProps } from "tailwind-variants";
import type { ChangeEvent } from "react";
import { Search } from "lucide-react";

const inputStyles = tv({
  base: "focus:outline-0 focus:outline-offset-0 focus:ring-0",
  variants: {
    intent: {
      primary: "input input-primary input-sm text-slate-600",
      valuePrimary: "input input-primary input-sm focus:outline-none",
      search: "input input-secondary !border-slate-200 !border",
      read: "input input-primary input-sm !bg-transparent",
    },
    fullWidth: {
      true: "w-full",
      false: "w-[24rem]",
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
});

interface Props extends VariantProps<typeof inputStyles> {
  defaultValue?: string;
  ariaLabel?: string;
  dataTestid?: string;
  required?: boolean;
  name?: string;
  onChange?: (e: ChangeEvent<HTMLInputElement>) => void;
}

function Input(props: Props) {
  return (
    <div>
      {props.intent === "primary" || props.intent === "valuePrimary" ? (
        <div>
          {props.intent === "primary" ? (
            <>
              <input
                required={props.required}
                type={props.type as string}
                name={props.name}
                className={inputStyles(props)}
                onChange={props.onChange}
                aria-label={props.ariaLabel}
                data-testid={props.dataTestid}
              />
            </>
          ) : (
            <>
              <input
                required={props.required}
                type={props.type as string}
                name={props.name}
                className={inputStyles(props)}
                onChange={props.onChange}
                aria-label={props.ariaLabel}
                data-testid={props.dataTestid}
                defaultValue={props.defaultValue}
              />
            </>
          )}
        </div>
      ) : (
        <>
          {props.intent === "read" ? (
            <input
              required={props.required}
              type={props.type as string}
              className={inputStyles(props)}
              onChange={props.onChange}
              aria-label={props.ariaLabel}
              data-testid={props.dataTestid}
              defaultValue={props.defaultValue}
              disabled
            />
          ) : (
            <div className={inputStyles(props)}>
              <Search className="w-4 h-4" />
              <input
                required={props.required}
                type={props.type as string}
                className="border-none focus:outline-none pl-2 w-full"
                onChange={props.onChange}
                aria-label={props.ariaLabel}
                data-testid={props.dataTestid}
              />
            </div>
          )}
        </>
      )}
    </div>
  );
}

export default Input;

import { tv, type VariantProps } from "tailwind-variants";
import { Fingerprint, Globe } from "lucide-react";

const tagStyles = tv({
  base: "badge badge-secondary badge-outline badge-lg rounded !text-slate-600",
  variants: {
    intent: {
      primary: "border-slate-100 pt-2 pb-2",
      landingpg: "bg-slate-100 text-slate-600 border-0 pt-2 pb-2",
      publicTag: "bg-slate-50 border-slate-100",
      privateTag: "bg-slate-50 border-slate-100",
    },
  },
  defaultVariants: {
    intent: "primary",
  },
});

interface Props extends VariantProps<typeof tagStyles> {
  [children: string]: any;
}

function Tag(props: Props) {
  return (
    <div className="pr-2">
      <div className={tagStyles(props)}>
        {props.intent === "publicTag" || props.intent === "privateTag" ? (
          <>
            {props.intent === "publicTag" ? (
              <>
                <Globe className="w-3 h-3" />
                <div className="pl-2 text-xs">{props.children}</div>
              </>
            ) : (
              <>
                <Fingerprint className="w-3 h-3" />
                <div className="pl-2 text-xs">{props.children}</div>
              </>
            )}
          </>
        ) : (
          <>
            {props.intent === "primary" ? (
              <div>{props.children}</div>
            ) : (
              <div className="flex justify-center items-center">
                <img
                  src="/imgs/landingPage/icons/ComingSoonIcon.svg"
                  alt="ComingSoon"
                  className="pr-2"
                />
                Coming soon...
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}

export default Tag;

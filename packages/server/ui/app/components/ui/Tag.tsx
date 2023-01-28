import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";
import { CircleDot, Database, File } from "lucide-react";

const tagStyles = cva("rounded flex items-center", {
  variants: {
    intent: {
      primary:
        "text-xs text-slate-600 w-fit h-fit bg-slate-50 px-2 py-1 border border-slate-100",
      landingpg:
        "text-xs text-slate-600 w-fit h-fit px-4 py-1 border border-slate-200",
      modelTag:
        "text-xs text-slate-600 w-fit h-fit bg-slate-50 px-2 py-1 border border-slate-100",
      datasetTag:
        "text-xs text-slate-600 w-fit h-fit bg-slate-50 px-2 py-1 border border-slate-100",
    },
  },
  defaultVariants: {
    intent: "primary",
  },
});

interface Props extends VariantProps<typeof tagStyles> {
  [children: string]: any;
}

function Tag({ intent, children }: Props) {
  return (
    <div className="pr-2">
      <div className={tagStyles({ intent })}>
        {intent === "modelTag" || intent === "datasetTag" ? (
          <>
            {intent === "modelTag" ? (
              <>
                <Database className="w-3 h-3" />
                <div className="pl-2">{children}</div>
              </>
            ) : (
              <>
                <File className="w-3 h-3" />
                <div className="pl-2">{children}</div>
              </>
            )}
          </>
        ) : (
          <>
            {intent === "primary" ? (
              <div>{children}</div>
            ) : (
              <div className="flex justify-center items-center">
                <img
                  src="/imgs/landingPage/ComingSoonIcon.svg"
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

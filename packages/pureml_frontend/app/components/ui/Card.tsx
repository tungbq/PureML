import { tv, type VariantProps } from "tailwind-variants";
import { Box, Database } from "lucide-react";

const cardStyles = tv({
  base: "rounded-lg border border-slate-200 justify-center px-6 py-4 font-normal focus:outline-none cursor-pointer hover:bg-slate-100 transform transition duration-300",
  variants: {
    intent: {
      modelCard: "text-slate-600",
      datasetCard: "text-slate-600",
    },
    fullWidth: {
      true: "w-full",
    },
  },
  defaultVariants: {
    intent: "modelCard",
    fullWidth: true,
  },
});
type CardProps = VariantProps<typeof cardStyles>;

interface Props extends CardProps {
  name: string;
  description: string;
  // tag1: string;
  tag2: any;
  onClick?: () => void;
}

export default function Card(props: Props) {
  return (
    <div onClick={props.onClick} className={cardStyles(props)}>
      <header className="pb-0 text-slate-600">
        <div className="flex items-center">
          {props.intent === "modelCard" ? (
            <Box className="text-slate-400 w-4" />
          ) : (
            <Database className="text-slate-400 w-4" />
          )}
          <span className="ml-2 truncate font-medium">{props.name}</span>
        </div>
      </header>
      <div className="text-sm font-normal truncate pt-2">
        {props.description}
      </div>
      <div className="flex pt-4">
        {/* {props.tag1} */}
        {props.tag2}
      </div>
    </div>
  );
}

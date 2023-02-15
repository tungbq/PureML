import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";
import { Box, Database } from "lucide-react";
import Tag from "./ui/Tag";

const cardStyles = cva(
  "items-center justify-center px-6 py-4 text-lg font-normal focus:outline-none cursor-pointer",
  {
    variants: {
      intent: {
        modelCard:
          "text-slate-600 rounded-lg border-2 border-slate-200 hover:bg-slate-50",
        datasetCard:
          "text-slate-600 rounded-lg border-2 border-slate-200 hover:bg-slate-50",
      },
      fullWidth: {
        true: "w-full",
      },
    },
    defaultVariants: {
      intent: "modelCard",
      fullWidth: true,
    },
  }
);
type CardProps = VariantProps<typeof cardStyles>;

interface Props extends CardProps {
  name: string;
  description: string;
  // tag1: string;
  tag2: string;
  onClick?: () => void;
}

export default function Card({
  intent,
  fullWidth,
  name,
  description,
  // tag1,
  tag2,
  onClick,
}: Props) {
  return (
    <div onClick={onClick} className={cardStyles({ intent, fullWidth })}>
      <header className="pb-0 text-slate-800">
        <div className="flex items-center">
          {intent === "modelCard" ? (
            <Box className="w-4" />
          ) : (
            <Database className="w-4" />
          )}
          <span className="ml-2 truncate text-sm font-medium">{name}</span>
        </div>
      </header>
      <div className="text-xs font-normal truncate pt-2">{description}</div>
      <div className="flex pt-4">
        {/* <Tag>{tag1}</Tag> */}
        {intent === "modelCard" ? (
          <Tag intent="modelTag">{tag2}</Tag>
        ) : (
          <Tag intent="datasetTag">{tag2}</Tag>
        )}
      </div>
    </div>
  );
}

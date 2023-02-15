import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";
import { CopyBlock, monoBlue } from "react-code-blocks";
import Link from "~/components/ui/Link";

const emptyStyles = cva("h-full flex justify-center items-center", {
  variants: {
    intent: {
      primary:
        "flex justify-center items-center bg-slate-150 drop-shadow-3xl rounded-lg",
    },
    fullWidth: {
      true: "w-full",
    },
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: true,
  },
});

const codeStyles = cva("h-fit ml-8 bg-slate-150 text-slate-400", {
  variants: {
    intent: {
      primary: "",
    },
    fullWidth: {
      true: "w-full",
    },
  },
  defaultVariants: {
    intent: "primary",
    fullWidth: true,
  },
});

interface Props extends VariantProps<typeof emptyStyles> {
  // [children: string]: any;
}

// export default function EmptyModel({ intent, fullWidth }: Props) {
//   return (
//     <div className="flex w-full justify-center items-center pt-24">
//       <div className="flex">
//         <div className="w-2/5">
//           <div className={emptyStyles({ intent, fullWidth })}>
//             <div className="px-4">
//               <div className="flex justify-center">
//                 <img
//                   src="/imgs/EmptyFolder.svg"
//                   alt="NoProjectsImg"
//                   width="40"
//                   height="40"
//                   className="w-10"
//                 />
//               </div>
//               <div className="flex justify-center mt-2 text-base text-slate-800">
//                 Register and Train Models
//               </div>
//               <div className="flex justify-center text-slate-400 text-sm mb-4">
//                 There are no Models added yet.
//               </div>
//               <div className="flex justify-center">
//                 <Link hyperlink="https://docs.pureml.com">Refer Docs</Link>
//               </div>
//             </div>
//           </div>
//         </div>
//         <div className="w-3/5">
//           <div className={codeStyles({ intent, fullWidth })}>
//             <div className="px-6 py-6 bg-slate-150">
//               <div className="text-sm text-slate-600 font-medium">
//                 Register new Model
//               </div>
//               <div className="text-zinc-400 text-sm">
//                 <CopyBlock
//                   text={`
// from pureml.decorators import model

// @model('iris_classifier')
// def train():
//     ...
//     return xgb_model
// `}
//                   language="python"
//                   theme={monoBlue}
//                   wrapLines={false}
//                 />
//               </div>
//             </div>
//           </div>
//         </div>
//       </div>
//     </div>
//   );
// }

export default function EmptyModel() {
  return (
    <div className="pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 min-w-72">
      <div className="rounded-lg border-2 border-slate-200 px-6 py-4">
        <div className="font-medium text-sm pb-6">There are no models yet</div>
        <div className="rounded-lg h-2 bg-slate-200 w-1/3" />
        <div className="pt-2"></div>
        <div className="rounded-lg h-2 bg-slate-200 w-2/3" />
      </div>
    </div>
  );
}

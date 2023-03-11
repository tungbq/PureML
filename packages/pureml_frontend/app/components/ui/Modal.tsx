import * as Dialog from "@radix-ui/react-dialog";
import { tv, type VariantProps } from "tailwind-variants";
import { X } from "lucide-react";
import type { ReactNode } from "react";

const editBtnStyles = tv({
  base: "flex items-center w-full px-4 py-2 font-medium focus:outline-none hover:bg-slate-200 hover:text-slate-600 rounded-lg",
  variants: {
    intent: {
      primary: "",
    },
  },
  defaultVariants: {
    intent: "primary",
  },
});

interface Props extends VariantProps<typeof editBtnStyles> {
  btnName: any;
  title: string;
  children: ReactNode;
}

export default function Modal(props: Props) {
  return (
    <div>
      <Dialog.Root>
        <Dialog.Trigger asChild>{props.btnName}</Dialog.Trigger>
        <Dialog.Portal>
          <Dialog.Overlay className="fixed absolute top-0 h-screen w-screen z-20 bg-zinc-800 opacity-60" />
          <div className="flex justify-center items-center h-full w-full z-50 fixed top-0 left-0">
            <Dialog.Content className=" bg-white rounded-md flex w-fit p-6 focus:outline-none absolute">
              <div className="w-full">
                <Dialog.Title className="text-base text-slate-800 font-medium">
                  {props.title}
                </Dialog.Title>
                <div className="pt-6 text-sm text-slate-600">
                  {props.children}
                </div>
                <Dialog.Close asChild>
                  <button
                    className="rounded-full h-6 w-6 flex items-center justify-content-center text-violet-400 absolute top-3.5 right-2.5"
                    aria-label="Close"
                  >
                    <X />
                  </button>
                </Dialog.Close>
              </div>
            </Dialog.Content>
          </div>
        </Dialog.Portal>
      </Dialog.Root>
    </div>
  );
}

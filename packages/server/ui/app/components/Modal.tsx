import * as Dialog from "@radix-ui/react-dialog";
import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";
import { X } from "lucide-react";
import Button from "./ui/Button";
import Input from "./ui/Input";

const editBtnStyles = cva(
  "flex items-center w-full px-4 py-2 font-medium focus:outline-none hover:bg-slate-200 hover:text-slate-600 rounded-lg",
  {
    variants: {
      intent: {
        editModal: "",
        successModal: "w-full",
        errorModal: "w-full",
      },
    },
    defaultVariants: {
      intent: "editModal",
    },
  }
);

interface Props extends VariantProps<typeof editBtnStyles> {
  btnName: string;
  title: string;
}

export default function Modal({ btnName, title, intent }: Props) {
  return (
    <div>
      <Dialog.Root>
        <Dialog.Trigger asChild>
          <button className={editBtnStyles({ intent })}>
            <div className="flex items-center">
              <div className="px-4">{btnName}</div>
            </div>
          </button>
        </Dialog.Trigger>
        <Dialog.Portal>
          <Dialog.Overlay className="fixed absolute top-0 h-screen w-screen z-20 bg-zinc-800 opacity-60" />
          <div className="flex justify-center items-center h-[36rem]">
            <Dialog.Content className="z-50 bg-white rounded-md flex w-96 p-6 focus:outline-none absolute">
              <div className="w-full">
                <Dialog.Title className="text-base text-slate-900">
                  {title}
                </Dialog.Title>
                {intent === "editModal" ? (
                  <>
                    <div className="pt-6">
                      <div>
                        <span className="text-sm text-slate-600">Name</span>
                        <Input placeholder="Enter your name" />
                      </div>
                      <div>
                        <div className="text-sm text-slate-600 pt-4">
                          Description
                        </div>
                        <textarea
                          className="textarea w-full bg-transparent text-sm border border-slate-600 rounded-md h-full hover:border-blue-750 focus:outline-none focus:border-blue-750 max-h-[200px] p-4"
                          placeholder="Add your Description"
                        />
                      </div>
                    </div>
                    <div className="pt-12 grid justify-items-end w-full">
                      <div className="flex justify-between w-1/2">
                        <Dialog.Close asChild>
                          <Button intent="secondary" icon="" fullWidth={false}>
                            Cancel
                          </Button>
                        </Dialog.Close>
                        <Button
                          intent="primary"
                          icon=""
                          fullWidth={false}
                          className="pl-2"
                        >
                          Save
                        </Button>
                      </div>
                    </div>
                  </>
                ) : (
                  <>
                    <div className="pt-6 text-sm text-slate-600">
                      {intent === "successModal" ? (
                        <>
                          <div>
                            Are you sure you want to archive this project?
                          </div>
                          <div>
                            This action can be undone by unarchiving this
                            project.
                          </div>
                        </>
                      ) : (
                        <>
                          <div>Are you sure you want to delete this model?</div>
                          <div>This action cant be undone.</div>
                        </>
                      )}
                    </div>
                    <div className="pt-12 grid justify-items-end w-full">
                      <div className="flex justify-between w-1/2">
                        <Dialog.Close asChild>
                          <Button intent="secondary" icon="" fullWidth={false}>
                            Cancel
                          </Button>
                        </Dialog.Close>
                        {intent === "successModal" ? (
                          <Button
                            intent="primary"
                            icon=""
                            fullWidth={false}
                            className="pl-2"
                          >
                            Archive
                          </Button>
                        ) : (
                          <Button
                            intent="danger"
                            icon=""
                            fullWidth={false}
                            className="!pl-2"
                          >
                            Delete
                          </Button>
                        )}
                      </div>
                    </div>
                  </>
                )}
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

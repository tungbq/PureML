import { useState } from "react";
import { Menu, X } from "lucide-react";
import Button from "../ui/Button";

export default function Navbar() {
  const [open, setOpen] = useState(true);
  if (!open)
    return (
      <div className="max-w-full sm:max-w-7xl sm:px-24">
        <div className="flex p-4 md:px-12 w-full justify-between max-w-7xl">
          <img src="/LogoWText.svg" alt="" width="96" height="96" />
          <X
            className="sm:hidden text-slate-900 cursor-pointer w-8 h-8"
            onClick={() => setOpen(!open)}
          />
        </div>
        <div className="flex flex-col gap-y-2 p-4 font-medium text-slate-850">
          <div className="flex items-center">
            <a href="https://docs.pureml.com" className="w-max">
              Docs
            </a>
          </div>
          <div className="flex items-center">
            <a href="https://discord.gg/xNUHt9yguJ" className="w-max">
              Join Discord
            </a>
          </div>
          <div className="flex items-center text-blue-600 hover:text-black pb-2 border-b border-slate-200">
            <a href="https://app.pureml.com">Sign in</a>
          </div>
        </div>
      </div>
    );
  return (
    <div className="bg-slate-50 flex justify-center items-center">
      <div className="flex p-4 w-full justify-between max-w-7xl">
        <img src="/LogoWText.svg" alt="" width="96" height="96" />
        <Menu
          className="sm:hidden text-slate-900 cursor-pointer w-8 h-8"
          onClick={() => setOpen(!open)}
        />
        <div className="hidden sm:flex flex font-medium">
          <div className="px-4 flex justify-center items-center">
            <a href="https://docs.pureml.com" className="w-full text-slate-850">
              Docs
            </a>
          </div>
          <div className="px-4 flex justify-center items-center">
            <a
              href="https://discord.gg/xNUHt9yguJ"
              className="w-max text-slate-850"
            >
              Join Discord
            </a>
          </div>
          <div className="px-4 flex justify-center items-center">
            <a
              className="github-button"
              href="https://github.com/pureml-inc/pureml"
              data-color-scheme="no-preference: dark_dimmed; light: light_high_contrast; dark: light;"
              data-size="large"
              data-show-count="true"
              aria-label="Star pureml-inc/pureml on GitHub"
            >
              Star
            </a>
          </div>
          <div className="px-4 flex justify-center items-center text-blue-600 hover:text-black">
            <a href="https://app.pureml.com">
              <Button intent="primary" icon="">
                Sign in
              </Button>
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}

import clsx from "clsx";
import { File, Box, Database } from "lucide-react";
import Input from "./ui/Input";
import Button from "./ui/Button";
import { Link, useMatches, useNavigate } from "@remix-run/react";
import type { VariantProps } from "class-variance-authority";
import { cva } from "class-variance-authority";
import AvatarIcon from "./ui/Avatar";
import Dropdown from "./ui/Dropdown";

function linkCss(currentPage: boolean) {
  return clsx(
    currentPage ? " text-blue-700 " : " text-slate-600 ",
    " hover:text-blue-700 flex justify-center items-center px-5 cursor-pointer "
  );
}

const navbarStyles = cva(
  "fixed z-20 h-18 px-12 py-4 w-full bg-slate-0 flex justify-between text-sm font-medium border-b-2 border-slate-200",
  {
    variants: {
      intent: {
        loggedIn: "",
        loggedOut: "",
      },
      fullWidth: {
        true: "w-full",
      },
    },
    defaultVariants: {
      intent: "loggedOut",
      fullWidth: true,
    },
  }
);
interface Props extends VariantProps<typeof navbarStyles> {
  [user: string]: any;
}

export default function NavBar({ intent, fullWidth, user }: Props) {
  const navigate = useNavigate();
  const matches = useMatches();
  const pathname = matches[1].pathname;

  return (
    <>
      <div className={navbarStyles({ intent, fullWidth })}>
        <div className="flex">
          <a href="/models" className="flex items-center justify-center pr-8">
            <img src="/LogoWText.svg" alt="Logo" width="140" height="96" />
          </a>
          <Input
            intent="search"
            placeholder="Search models, datasets, users..."
            fullWidth={false}
          />
        </div>

        <div className="flex justify-center items-center">
          <div
            onClick={() => {
              navigate(`/models`);
            }}
            className={`${linkCss(pathname === `/models`)}`}
          >
            <Box className="w-4 h-4" />
            <span className="pl-2">Models</span>
          </div>

          <div
            onClick={() => {
              navigate(`/datasets`);
            }}
            className={`${linkCss(pathname === `/datasets`)}`}
          >
            <Database className="w-4 h-4" />
            <span className="pl-2">Datasets</span>
          </div>

          <a
            href="https://docs.pureml.com"
            className="flex justify-center items-center cursor-pointer px-5 hover:text-blue-700 border-r-2 border-slate-slate-200 font-medium text-slate-600"
          >
            <File className="w-4 h-4" />
            <span className="pl-2">Docs</span>
          </a>
          {intent === "loggedOut" ? (
            <>
              <div className="w-full flex justify-center items-center px-5">
                <a href="/auth/signin">Sign in</a>
              </div>
              <Button intent="primary" icon="">
                Sign up
              </Button>
            </>
          ) : (
            <div className="w-full flex justify-center items-center px-5">
              <Dropdown intent="primary">
                <AvatarIcon>{user}</AvatarIcon>
              </Dropdown>
            </div>
          )}
        </div>
      </div>
    </>
  );
}

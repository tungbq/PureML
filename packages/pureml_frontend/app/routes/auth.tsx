import type { MetaFunction } from "@remix-run/node";
import { Outlet } from "@remix-run/react";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Join Us | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export default function AuthLayout() {
  return (
    <div className="w-screen h-screen flex items-center justify-center bg-slate-0">
      <div className="g-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white">
        <div className="w-fit text-center">
          <div className="flex justify-center items-center pb-16">
            <img src="/PureMLLogoWText.svg" alt="Logo" className="w-36" />
          </div>
          <Outlet />
        </div>
      </div>
    </div>
  );
}

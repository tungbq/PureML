import type { MetaFunction } from "@remix-run/node";
import { Meta, Outlet } from "@remix-run/react";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Join Us | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export default function AuthLayout() {
  return (
    <main className="w-screen h-screen flex justify-center items-center bg-slate-0">
      <head>
        <Meta />
      </head>
      <div className="hidden h-screen md:w-2/5 bg-slate-50 md:flex md:flex-col md:justify-center 4xl:items-center md:pl-24 xl:pl-44 md:overflow-visible">
        <img src="/Logo.svg" alt="logo" className="w-24" />
        <h1 className="text-slate-600 text-5xl font-medium mt-6">
          Welcome to{" "}
          <span className="font-medium text-slate-900 mt-11">PureML</span>
        </h1>
        <img
          src="/imgs/AuthCodeSnippet.svg"
          alt="SignInCode"
          className="md:w-80 lg:w-96 xl:w-[30rem] 2xl:w-[42rem] 4xl:w-[54rem] max-w-[710px] mt-12 -ml-4 z-10"
        />
      </div>
      <Outlet />
    </main>
  );
}

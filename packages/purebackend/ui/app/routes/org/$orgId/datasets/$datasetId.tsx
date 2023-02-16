import type { MetaFunction } from "@remix-run/node";
import { Meta, Outlet } from "@remix-run/react";
import Breadcrumbs from "~/components/Breadcrumbs";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Dataset Details | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export default function DatasetIndex() {
  return (
    <div id="datasets">
      <head>
        <Meta />
      </head>
      <div className="flex flex-col">
        <div className="px-12 sticky top-16 bg-slate-0 w-full z-10">
          <Breadcrumbs />
        </div>
        <Outlet />
      </div>
    </div>
  );
}

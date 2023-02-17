import type { MetaFunction } from "@remix-run/node";
import { Meta, Outlet } from "@remix-run/react";
import Breadcrumbs from "~/components/Breadcrumbs";
import Tag from "~/components/ui/Tag";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Model Details | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export default function ModelIndex() {
  return (
    <div id="models">
      <head>
        <Meta />
      </head>
      <div className="flex flex-col">
        <div className="px-12 sticky top-16 bg-slate-0 w-full z-10">
          <Breadcrumbs />
          <div className="flex pt-6 py-4">
            <Tag>Dummy Tag 1</Tag>
            <Tag>Tag 2</Tag>
            <Tag>Dummy</Tag>
          </div>
        </div>
        <Outlet />
      </div>
    </div>
  );
}

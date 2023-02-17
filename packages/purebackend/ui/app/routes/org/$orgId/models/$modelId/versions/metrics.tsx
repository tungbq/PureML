import { Outlet } from "@remix-run/react";
import Tabbar from "~/components/Tabbar";

export default function Metrics() {
  return (
    <div className="">
      <Tabbar intent="primaryModelTab" tab="versions" />
      <Outlet />
    </div>
  );
}

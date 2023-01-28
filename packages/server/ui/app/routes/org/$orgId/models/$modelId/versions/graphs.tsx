import { Outlet } from "@remix-run/react";
import Tabbar from "~/components/Tabbar";

export default function Graphs() {
  return (
    <div className="">
      <Tabbar intent="primaryModelTab" tab="versions" />
      <Outlet />
    </div>
  );
}

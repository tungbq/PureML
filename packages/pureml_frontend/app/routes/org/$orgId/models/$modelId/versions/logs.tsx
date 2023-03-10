import { Outlet } from "@remix-run/react";

export default function Metrics() {
  return (
    <div className="h-full">
      <Outlet />
    </div>
  );
}

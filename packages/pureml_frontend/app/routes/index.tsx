import { Outlet } from "@remix-run/react";
import { Suspense } from "react";
import Loader from "~/components/ui/Loading";

export default function Index() {
  return (
    <Suspense fallback={<Loader />}>
      <Outlet />
    </Suspense>
  );
}

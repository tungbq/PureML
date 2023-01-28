import { Outlet, useLoaderData } from "@remix-run/react";
import NavBar from "~/components/Navbar";
import { getSession } from "~/session";
import { fetchUserSettings } from "./api/auth.server";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const profile = await fetchUserSettings(accesstoken);
  return profile;
}

export default function OrgLayout() {
  const prof = useLoaderData();
  return (
    <div>
      {prof ? (
        <NavBar intent="loggedIn" user={prof[0].name.charAt(0).toUpperCase()} />
      ) : (
        <NavBar intent="loggedOut" />
      )}
      <div className="pt-16 h-screen w-screen">
        <Outlet />
      </div>
    </div>
  );
}

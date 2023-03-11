import { Outlet, useLoaderData } from "@remix-run/react";
import NavBar from "~/components/Navbar";
import { getSession } from "~/session";
import { fetchUserSettings } from "./api/auth.server";
import type { MetaFunction } from "@remix-run/node";
import { fetchOrgDetails } from "./api/org.server";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Settings | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const profile = await fetchUserSettings(accesstoken);
  const orgId = session.get("orgId");
  const orgDetails = await fetchOrgDetails(orgId, session.get("accessToken"));
  return { profile, orgDetails };
}

export default function SettingsLayout() {
  const prof = useLoaderData();
  return (
    <div className="h-screen overflow-hidden">
      {prof ? (
        <NavBar
          intent="loggedIn"
          user={prof.profile[0].name.charAt(0).toUpperCase()}
          orgName={
            <a
              href={`/org/${prof.orgDetails[0].name}`}
              className="flex items-center justify-center"
            >
              {prof.orgDetails[0].name}
            </a>
          }
          orgAvatarName={prof.orgDetails[0].name.charAt(0)}
        />
      ) : (
        <NavBar
          intent="loggedOut"
          user=""
          orgName={
            <a href="/models" className="flex items-center justify-center pr-8">
              <img src="/PureMLLogoWText.svg" alt="Logo" className="w-20" />
            </a>
          }
          orgAvatarName=""
        />
      )}
      <div className="pb-12 pt-6 h-full bg-slate-50">
        <div className="flex justify-center">
          <div className="sticky top-0 px-12 flex justify-between font-medium text-slate-800 w-full max-w-screen-2xl">
            Settings
          </div>
        </div>
        <Outlet />
      </div>
    </div>
  );
}

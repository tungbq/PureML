import type { MetaFunction } from "@remix-run/node";
import { Outlet, useLoaderData } from "@remix-run/react";
import NavBar from "~/components/Navbar";
import { getSession } from "~/session";
import { fetchUserSettings } from "./api/auth.server";
import { fetchOrgDetails } from "./api/org.server";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Models | PureML",
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

export default function ModelsLayout() {
  const prof = useLoaderData();
  return (
    <div className="bg-slate-50 flex flex-col h-screen overflow-hidden">
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
      <div className="flex justify-center w-full">
        <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
          <div className="flex justify-between px-12 pt-6 pb-4 font-medium text-slate-800">
            Models
          </div>
          <div className="h-[95%] overflow-auto pb-8">
            <Outlet />
          </div>
        </div>
      </div>
    </div>
  );
}

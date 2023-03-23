import type { ActionArgs, MetaFunction } from "@remix-run/node";
import { redirect } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";
import { toast } from "react-toastify";
import NavBar from "~/components/Navbar";
import AvatarIcon from "~/components/ui/Avatar";
import Button from "~/components/ui/Button";
import Error404 from "~/Error404";
import { commitSession, getSession } from "~/session";
import { fetchUserSettings } from "./api/auth.server";
import { fetchAllOrgs, fetchOrgDetails } from "./api/org.server";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Switch Organization | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const orgId = session.get("orgId");
  if (!orgId) return null;
  const profile = await fetchUserSettings(accesstoken);
  const orgDetails = await fetchOrgDetails(orgId, session.get("accessToken"));
  const allOrgs = await fetchAllOrgs(accesstoken);
  return { session, profile, orgDetails, allOrgs };
}

export async function action({ request }: ActionArgs) {
  const session = await getSession(request.headers.get("Cookie"));
  const form = await request.formData();
  const orgID = form.get("orgid");
  const orgName = form.get("orgname");
  session.unset("orgId");
  session.unset("orgName");
  session.set("orgId", orgID);
  session.set("orgName", orgName);
  return redirect("/models", {
    headers: { "Set-Cookie": await commitSession(session) },
  });
}

export default function SwitchOrg() {
  const data = useLoaderData();
  if (!data) return <Error404 />;
  return (
    <div className="bg-slate-50 flex flex-col h-screen overflow-hidden">
      {data ? (
        <NavBar
          intent="loggedIn"
          user={data.profile[0].name.charAt(0).toUpperCase()}
          orgName={
            <a
              href={`/org/${data.orgDetails[0].name}`}
              className="flex items-center justify-center"
            >
              {data.orgDetails[0].name}
            </a>
          }
          orgAvatarName={data.orgDetails[0].name.charAt(0)}
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
          <div className="px-12 pt-6 pb-6 h-[90%] overflow-auto">
            <div className="text-slate-600 font-medium">
              Switch Organization
            </div>
            <div className="mt-6">
              {data.allOrgs[0] ? (
                <>
                  {data.allOrgs.map((org: any, index: number) => (
                    <div className="pb-6" key={index}>
                      <Form
                        method="post"
                        className="text-sm text-slate-600 font-medium"
                      >
                        <div className="bg-slate-100 rounded-2xl p-4 flex items-center justify-between w-2/3 gap-x-4">
                          <div className="flex justify-start w-4/5">
                            <div className="flex justify-center items-center">
                              <AvatarIcon intent="org">
                                {org.org.name.charAt(0).toUpperCase()}
                              </AvatarIcon>
                            </div>
                            <div className="flex flex-col justify-center px-4 text-slate-600">
                              <input
                                className="hidden"
                                type="text"
                                name="orgid"
                                defaultValue={org.org.uuid}
                              />
                              <input
                                className="hidden"
                                type="text"
                                name="orgname"
                                defaultValue={org.org.name}
                              />
                              <div className="font-medium">{org.org.name}</div>
                              <div className="font-normal">
                                {org.org.description || "Sample description"}
                              </div>
                            </div>
                          </div>
                          <div className="w-fit">
                            <Button
                              intent="secondary"
                              fullWidth={false}
                              type="submit"
                              onClick={() => {
                                toast.success("Switching organization!");
                              }}
                            >
                              Switch
                            </Button>
                          </div>
                        </div>
                      </Form>
                    </div>
                  ))}
                </>
              ) : (
                "No organizations found :("
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

// ############################ error boundary ###########################

export function ErrorBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center bg-slate-50">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/ErrorFunction.gif" alt="Error" width="500" />
    </div>
  );
}

export function CatchBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center bg-slate-50">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/ErrorFunction.gif" alt="Error" width="500" />
    </div>
  );
}

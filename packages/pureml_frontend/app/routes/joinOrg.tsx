import { json, MetaFunction } from "@remix-run/node";
import { redirect } from "@remix-run/node";
import {
  Form,
  useActionData,
  useLoaderData,
  useNavigate,
} from "@remix-run/react";
import { X } from "lucide-react";
import { useEffect, useState } from "react";
import { toast } from "react-toastify";
import NavBar from "~/components/Navbar";
import Button from "~/components/ui/Button";
import Input from "~/components/ui/Input";
import { getSession } from "~/session";
import { fetchUserSettings } from "./api/auth.server";
import { fetchAllOrgs, fetchJoinOrg, fetchOrgDetails } from "./api/org.server";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Join Organization | PureML",
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

export const action = async ({ request }: any) => {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const form = await request.formData();
  const orgJoinCode = form.get("orgcode");
  const data = await fetchJoinOrg(orgJoinCode, accesstoken);
  if (data.statusText === "OK") return data;
  else if (data.statusText === "Conflict") return data.statusText;
  else if (data.statusText === "Not found") return data.statusText;
  else return null;
};

export default function JoinOrg() {
  const prof = useLoaderData();
  const actData = useActionData();
  const navigate = useNavigate();
  const [actionData, setActionData] = useState(false);
  useEffect(() => {
    if (actData) {
      setActionData(true);
      toast.success("Organization joined!");
      navigate("/switchOrg");
    } else if (actData === "Not found") {
      toast.error("Oops! Joining code not found!");
    } else if (actData === "Conflict") {
      toast.error("Organization already exists!");
    } else if (actData === null) {
      toast.error("Oops! Something went wrong!");
    }
  }, [actionData]);
  return (
    <div className="flex flex-col w-screen h-screen overflow-hidden">
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
      <div className="bg-slate-50 w-full h-full flex justify-center items-center">
        <div className="bg-slate-0 border border-slate-200 p-4 rounded-lg h-fit w-fit z-50">
          <div className="flex justify-between text-slate-800 font-medium pb-8">
            Join Organization
            <X
              className="text-slate-400 w-4 cursor-pointer"
              onClick={() => {
                navigate("/models");
              }}
            />
          </div>
          <Form method="post" reloadDocument className="text-sm text-slate-600">
            <div className="pb-6">
              <label htmlFor="orgcode">
                <div className="pb-2 font-medium">Organization Code</div>
                <Input
                  intent="primary"
                  type="text"
                  name="orgcode"
                  fullWidth={false}
                  aria-label="orgcode"
                  data-testid="orgcode"
                  required
                />
              </label>
            </div>
            <div className="w-full">
              <Button intent="primary" type="submit">
                Join Now
              </Button>
            </div>
          </Form>
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
